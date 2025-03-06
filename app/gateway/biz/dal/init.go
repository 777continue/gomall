package dal

import (
	"github.com/777continue/gomall/app/frontend/biz/dal/mysql"
	"github.com/777continue/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
