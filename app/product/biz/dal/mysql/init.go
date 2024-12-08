// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"os"

	"github.com/777continue/gomall/app/product/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	test_dsn := "root:root@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(test_dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		klog.Debugf("need demo data = %v", needDemoData)
		if needDemoData {
			DB.AutoMigrate( //nolint:errcheck
				&model.Product{},
				&model.Category{},
			)
			klog.Debug("insert sample data")
			InsertSample(DB)
		}
	}
}

func InsertSample(DB *gorm.DB) {
	categories := []model.Category{
		{Model: gorm.Model{ID: 1}, Name: "snacks", Description: "snackss"},
		{Model: gorm.Model{ID: 2}, Name: "drinks", Description: "drinks"},
	}
	DB.Create(&categories)
	products := []model.Product{
		{Model: gorm.Model{ID: 1}, Name: "candy", Description: "", Picture: "/static/image/candy.jpg", Price: 5.00},
		{Model: gorm.Model{ID: 2}, Name: "chips", Description: "", Picture: "/static/image/chips.jpg", Price: 8.00},
		{Model: gorm.Model{ID: 3}, Name: "coke", Description: "", Picture: "/static/image/coke.jpeg", Price: 3.50},
		{Model: gorm.Model{ID: 4}, Name: "coke2", Description: "", Picture: "/static/image/coke2.jpeg", Price: 3.50},
		{Model: gorm.Model{ID: 5}, Name: "latiao", Description: "", Picture: "/static/image/latiao.jpeg", Price: 6.00},
		{Model: gorm.Model{ID: 6}, Name: "nut", Description: "", Picture: "/static/image/nut.jpeg", Price: 8.00},
	}
	DB.Create(&products)
	DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 1 ), ( 2, 1 ), ( 3, 2 ), ( 4, 2 ), ( 5, 1 ), ( 6, 1 )")
}
