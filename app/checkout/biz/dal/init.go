package dal

import (
	"github.com/777continue/gomall/app/checkout/biz/dal/mysql"
	"github.com/777continue/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
