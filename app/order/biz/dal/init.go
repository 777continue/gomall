package dal

import (
	"github.com/777continue/gomall/app/order/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
