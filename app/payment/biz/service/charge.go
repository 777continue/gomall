package service

import (
	"context"
	"time"

	"github.com/777continue/gomall/app/payment/biz/dal/mysql"
	"github.com/777continue/gomall/app/payment/biz/model"
	"github.com/777continue/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// 初始化随机数种子
	rand.Seed(uint64(time.Now().UnixNano()))

	// 生成随机状态
	status := "success"
	if rand.Intn(100) >= 60 {
		status = "fail"
	}

	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	if status == "fail" {
		return &payment.ChargeResp{
			TransactionId: transactionId.String(),
			Status:        status, // 返回支付状态
		}, nil
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}

	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
		Status:        status, // 返回支付状态
	}, nil
}
