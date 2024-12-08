package dal

import (
	"github.com/777continue/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
