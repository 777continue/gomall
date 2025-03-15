package service

import (
	"context"
	"time"

	common "github.com/777continue/gomall/app/frontend/hertz_gen/frontend/common"
	frontendutils "github.com/777continue/gomall/app/frontend/utils"
	rpcorder "github.com/777continue/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	order_client "github.com/777continue/gomall/rpc_gen/rpc/order"
	product_client "github.com/777continue/gomall/rpc_gen/rpc/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type Consignee struct {
	Email string
	Addr  string
	Name  string
}

type Order struct {
	Consignee   Consignee
	OrderId     string
	CreatedDate string
	OrderState  string
	Cost        float32
	Items       []OrderItem
}

type OrderItem struct {
	ProductId   uint32
	ProductName string
	Picture     string
	Qty         uint32
	Cost        float32
}
type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	var orders []*Order
	listOrderResp, err := order_client.Client.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	if listOrderResp == nil || len(listOrderResp.Orders) == 0 {
		return utils.H{
			"title":  "Order",
			"orders": orders,
		}, nil
	}

	for _, v := range listOrderResp.Orders {
		var items []OrderItem
		var total float32
		if len(v.OrderItems) > 0 {
			for _, vv := range v.OrderItems {
				total += vv.Cost
				i := vv.Item
				productResp, err := product_client.Client.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId})
				if err != nil {
					return nil, err
				}
				if productResp.Product == nil {
					continue
				}
				p := productResp.Product
				items = append(items, OrderItem{
					ProductId:   i.ProductId,
					Qty:         uint32(i.Quantity),
					ProductName: p.Name,
					Picture:     p.Picture,
					Cost:        vv.Cost,
				})
			}
		}
		timeObj := time.Unix(int64(v.CreatedAt), 0)
		orders = append(orders, &Order{
			Cost:        total,
			Items:       items,
			CreatedDate: timeObj.Format("2006-01-02 15:04:05"),
			OrderId:     v.OrderId,
			Consignee:   Consignee{Addr: v.Address},
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": orders,
	}, nil
}
