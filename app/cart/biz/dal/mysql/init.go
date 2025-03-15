package mysql

import (
	"fmt"
	"os"
	"sync"

	"github.com/777continue/gomall/app/cart/biz/model"
	"github.com/777continue/gomall/app/cart/conf"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	err  error
	once sync.Once
)

func Init() {
	once.Do(func() {
		klog.Info("Initializing MySQL connection...")
		dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
		DB, err = gorm.Open(mysql.Open(dsn),
			&gorm.Config{
				PrepareStmt:            true,
				SkipDefaultTransaction: true,
			},
		)
		if err != nil {
			klog.Errorf("Failed to connect to MySQL: %v", err)
			panic(err)
		}

		err = DB.AutoMigrate(&model.Cart{})
		if err != nil {
			klog.Errorf("Failed to auto migrate Cart model: %v", err)
			panic(err)
		}
		klog.Info("MySQL connection initialized successfully")
	})
}
