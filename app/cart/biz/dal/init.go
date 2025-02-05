package dal

import (
	"github.com/777continue/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
