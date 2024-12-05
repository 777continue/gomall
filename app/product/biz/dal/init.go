package dal

import (
	"github.com/777continue/gomall/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
