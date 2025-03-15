package service

import (
	"context"
	"fmt"

	"github.com/777continue/gomall/app/order/biz/dal/mysql"
	"github.com/777continue/gomall/app/order/biz/model"
	order "github.com/777continue/gomall/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	if len(req.OrderItems) == 0 {
		err = fmt.Errorf("OrderItems empty")
		return
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()

		o := &model.Order{
			OrderId:    orderId.String(),
			OrderState: model.OrderStatePlaced,
			UserId:     req.UserId,
			Consignee: model.Consignee{
				Name: req.Name,
				Addr: req.Address,
			},
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		var itemList []*model.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &model.OrderItem{
				OrderIdRefer: o.OrderId,
				ProductId:    v.Item.ProductId,
				Quantity:     int32(v.Item.Quantity),
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			OrderId: orderId.String(),
		}

		return nil
	})

	return
}
