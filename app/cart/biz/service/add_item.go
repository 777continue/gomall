package service

import (
	"context"

	"github.com/777continue/gomall/app/cart/biz/dal/mysql"
	"github.com/777continue/gomall/app/cart/biz/model"
	cart "github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	klog.Infof("add cart item: %v", req)
	productResp, err := product_client.Client.GetProduct(s.ctx, &rpcproduct.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	if mysql.DB == nil {
		klog.Error("db is nil")
	} else {
		klog.Info("db is not nil")
	}
	err = model.AddCart(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	klog.Info("add cart item success")
	return &cart.AddItemResp{}, nil
}
