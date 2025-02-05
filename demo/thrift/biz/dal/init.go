package dal

import (
	"github.com/777continue/gomall/demo/thrift/biz/dal/mysql"
	//"github.com/777continue/gomall/demo/thrift/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
