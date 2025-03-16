package service

import (
	"context"
	"testing"

	"github.com/777continue/gomall/app/checkout/biz/dal/redis"
	checkout "github.com/777continue/gomall/rpc_gen/kitex_gen/checkout"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	order_client "github.com/777continue/gomall/rpc_gen/rpc/order"
	payment_client "github.com/777continue/gomall/rpc_gen/rpc/payment"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()
	redis.Init()
	payment_client.DefaultClient()
	product_client.DefaultClient()
	cart_client.DefaultClient()
	order_client.DefaultClient()
	s := NewCheckoutService(ctx)
	// init req and assert value
	req := &checkout.CheckoutReq{
		UserId:  2,
		Address: "123 Main St",
	}
	resp, err := s.Run(req)
	klog.Errorf("err: %v", err)
	klog.Infof("resp: %v", resp)

	// todo: edit your unit test

}
