package service

import (
	"context"
	"testing"

	"github.com/777continue/gomall/app/checkout/biz/dal/redis"
	checkout "github.com/777continue/gomall/rpc_gen/kitex_gen/checkout"
)

func TestCheckout_Run(t *testing.T) {
	ctx := context.Background()
	redis.Init()
	s := NewCheckoutService(ctx)
	// init req and assert value
	req := &checkout.CheckoutReq{
		UserId:  2,
		Address: "123 Main St",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
