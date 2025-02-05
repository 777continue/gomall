package dal

import (
	"github.com/777continue/gomall/app/email/biz/dal/mysql"
	"github.com/777continue/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
