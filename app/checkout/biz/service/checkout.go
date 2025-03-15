package service

import (
	"context"
	"fmt"
	"time"

	"github.com/777continue/gomall/app/checkout/biz/dal/redis"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/777continue/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/order"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/payment"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	order_client "github.com/777continue/gomall/rpc_gen/rpc/order"
	payment_client "github.com/777continue/gomall/rpc_gen/rpc/payment"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
}

func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

type stockDec struct {
	productId uint32
	num       uint32
	lockKey   string
	lockValue string
}
type checkoutProcess struct {
	total      float32
	orderId    string
	orderItems []*order.OrderItem
	stockDec   []stockDec
}

func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// get cart
	cartResult, err := s.getCart(req.UserId)
	if err != nil {
		return nil, err
	}

	// Distributed Transaction
	process := &checkoutProcess{}
	if err := s.ExecTransaction(req, cartResult, process); err != nil {
		return nil, err
	}
	// call payment
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: process.orderId,
		Amount:  process.total,
	}
	_, err = payment_client.Client.Charge(s.ctx, payReq)

	if err != nil {
		return nil, kerrors.NewBizStatusError(5004005, err.Error())
	}

	// TODO: impl GetUserEmail api

	/*email, err := user_client.Client.GetUserEmail(s.ctx, &user.GetUserEmailReq{UserId: req.UserId})
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)*/

	return &checkout.CheckoutResp{
		OrderId:       process.orderId,
		TransactionId: process.orderId, // Assuming paymentResult.TransactionId is same as orderId
	}, nil
}

func (s *CheckoutService) getCart(userId uint32) (*cart.GetCartResp, error) {
	cartResult, err := cart_client.Client.GetCart(s.ctx, &cart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}
	return cartResult, nil
}

func (s *CheckoutService) ExecTransaction(req *checkout.CheckoutReq, cartResult *cart.GetCartResp, process *checkoutProcess) error {
	var err error

	// 先执行库存扣减
	process.total, process.orderItems, err = s.decStock(cartResult, process)
	if err != nil {
		return kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	// 然后执行订单创建
	orderResp, err := order_client.Client.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:     req.UserId,
		OrderItems: process.orderItems,
	})
	if orderResp != nil {
		process.orderId = orderResp.OrderId
	}
	if err != nil {
		// 回滚
		if err := s.rollback(req, process); err != nil {
			return kerrors.NewGRPCBizStatusError(5004003, err.Error())
		}
		return kerrors.NewGRPCBizStatusError(5004004, err.Error())
	}

	// 提交
	if err := s.commit(req, process); err != nil {
		return kerrors.NewGRPCBizStatusError(5004005, err.Error())
	}
	return nil
}

func (s *CheckoutService) decStock(cartResult *cart.GetCartResp, process *checkoutProcess) (float32, []*order.OrderItem, error) {
	var total float32
	var orderItems []*order.OrderItem

	for _, cartItem := range cartResult.Items {
		// 加锁
		lockKey := fmt.Sprintf("product_lock:%d", cartItem.ProductId)
		lockValue := fmt.Sprintf("%d", time.Now().UnixNano())
		locked, err := redis.RedisClient.SetNX(s.ctx, lockKey, lockValue, 30*time.Second).Result()
		if err != nil || !locked {
			return 0, nil, kerrors.NewGRPCBizStatusError(5004007, "failed to acquire lock")
		}
		//
		productResp, err := product_client.Client.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if err != nil {
			return 0, nil, err
		}
		if productResp.Product == nil {
			continue
		}
		if productResp.Product.Stock < cartItem.Quantity {
			return 0, nil, kerrors.NewGRPCBizStatusError(5004006,
				fmt.Sprintf("product %d stock is not enough, available: %d, required: %d",
					cartItem.ProductId, productResp.Product.Stock, cartItem.Quantity))
		}
		cost := productResp.Product.Price * float32(cartItem.Quantity)
		total += cost

		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
		process.stockDec = append(process.stockDec, stockDec{
			productId: cartItem.ProductId,
			num:       -cartItem.Quantity,
			lockKey:   lockKey,
			lockValue: lockValue,
		})

	}
	return total, orderItems, nil
}

func (s *CheckoutService) commit(req *checkout.CheckoutReq, process *checkoutProcess) error {
	if _, err := cart_client.Client.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId}); err != nil {
		klog.Error(err.Error())
		return err
	}
	for _, stockDec := range process.stockDec {
		if _, err := product_client.Client.UpdateProduct(s.ctx, &product.UpdateProductReq{
			Id:    stockDec.productId,
			Stock: -stockDec.num,
		}); err != nil {
			return err
		}
		if val, err := redis.RedisClient.Get(s.ctx, stockDec.lockKey).Result(); err == nil && val == stockDec.lockValue {
			redis.RedisClient.Del(s.ctx, stockDec.lockKey)
		}
	}
	return nil
}

func (s *CheckoutService) rollback(req *checkout.CheckoutReq, process *checkoutProcess) error {
	// TODO:  cancel order
	for _, stockDec := range process.stockDec {
		if val, err := redis.RedisClient.Get(s.ctx, stockDec.lockKey).Result(); err == nil && val == stockDec.lockValue {
			redis.RedisClient.Del(s.ctx, stockDec.lockKey)
		}
	}
	return nil
}
