package service

import (
	"context"
	"strconv"

	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	frontendUtils "github.com/777continue/gomall/app/frontend/utils"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/cart"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	cart_client "github.com/777continue/gomall/rpc_gen/rpc/cart"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	cartResp, err := cart_client.Client.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := product_client.Client.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(p.Price) * float64(item.Quantity)

	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}