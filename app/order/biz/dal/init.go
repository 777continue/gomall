package dal

import (
	"github.com/777continue/gomall/app/order/biz/dal/mysql"
	"github.com/777continue/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
