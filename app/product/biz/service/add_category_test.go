package service

import (
	"context"
	"testing"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
)

func TestAddCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddCategoryService(ctx)
	// init req and assert value

	req := &product.AddCategoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
