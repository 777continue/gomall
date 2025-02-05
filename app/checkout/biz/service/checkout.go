package service

import (
	"context"
	"strconv"

	"github.com/777continue/gomall/app/checkout/infra/mq"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/777continue/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/email"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/order"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/payment"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	order_client "github.com/777continue/gomall/rpc_gen/rpc/order"
	payment_client "github.com/777continue/gomall/rpc_gen/rpc/payment"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := cart_client.Client.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)

	for _, cartItem := range cartResult.Items {
		productResp, resultErr := product_client.Client.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	var orderId string
	zipcode, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := order_client.Client.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(zipcode),
		},
		OrderItems: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = cart_client.Client.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := payment_client.Client.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
