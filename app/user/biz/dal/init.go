package dal

import (
	"github.com/777continue/gomall/app/user/biz/dal/mysql"
	"github.com/777continue/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}