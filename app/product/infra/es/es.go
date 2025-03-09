package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/777continue/gomall/app/product/biz/dal/mysql"
	"github.com/777continue/gomall/app/product/biz/model"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	Client *elasticsearch.Client
)

func Init() error {
	// 创建 Elasticsearch 客户端
	var err error
	Client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",              // 添加用户名
		Password: "nilYY1ISoSO42u2Fofme", // 添加密码
	})
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}

	// 从数据库获取所有商品数据
	productQuery := model.NewProductQuery(context.Background(), mysql.DB)
	products, err := productQuery.SearchProducts("")
	if err != nil {
		log.Fatalf("Failed to get products from DB: %v", err)
	}

	// 同步数据到 Elasticsearch
	for _, p := range products {
		err := IndexProduct(context.Background(), "products", fmt.Sprintf("%d", p.ID), p)
		if err != nil {
			return err
		}
	}
	return nil
}

func IndexProduct(ctx context.Context, index string, productID string, document interface{}) error {
	// 将document转换为map[string]interface{}
	productMap := map[string]interface{}{
		"id":          document.(*model.Product).ID, // 修改为 *model.Product
		"name":        document.(*model.Product).Name,
		"description": document.(*model.Product).Description,
		"picture":     document.(*model.Product).Picture,
		"price":       document.(*model.Product).Price,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(productMap); err != nil {
		return err
	}

	// 创建 Elasticsearch 的索引请求
	req := esapi.IndexRequest{
		Index:      index,     // 指定索引名称
		DocumentID: productID, // 指定文档 ID
		Body:       &buf,      // 设置请求体为编码后的数据
		Refresh:    "true",    // 设置刷新策略，确保数据立即可搜索
	}

	// 执行请求
	res, err := req.Do(ctx, Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document: %s", res.String())
		return err
	}

	return nil
}

func SearchProducts(ctx context.Context, index string, query map[string]interface{}) (*esapi.Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := Client.Search(
		Client.Search.WithContext(ctx),
		Client.Search.WithIndex(index),         // products
		Client.Search.WithBody(&buf),           // query -json-> buf
		Client.Search.WithTrackTotalHits(true), //置搜索请求的 track_total_hits 参数为 true，表示返回匹配的文档总数。
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
