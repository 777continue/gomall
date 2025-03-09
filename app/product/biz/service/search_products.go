package service

import (
	"context"
	"encoding/json"

	"github.com/777continue/gomall/app/product/infra/es"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	log.Infof("SearchProductsService Run: %+v", req)
	// 定义 Elasticsearch 查询条件
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": req.Query,
			},
		},
	}
	klog.Infof("Elasticsearch 查询条件: %+v", query)
	// 从 Elasticsearch 获取数据
	res, err := es.SearchProducts(s.ctx, "products", query)
	if err != nil {
		klog.Errorf("搜索商品失败: %v", err)
		return nil, err
	}
	klog.Infof("Elasticsearch 返回的原始响应: %+v", res) // 解析 Elasticsearch 返回的结果
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	klog.Infof("解析后的 Elasticsearch 结果: %+v", result)
	// 将结果转换为所需的格式
	var results []*product.Product
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		results = append(results, &product.Product{
			Id:          uint32(source["id"].(float64)),
			Name:        source["name"].(string),
			Description: source["description"].(string),
			Picture:     source["picture"].(string),
			Price:       float32(source["price"].(float64)),
		})
	}

	return &product.SearchProductsResp{Results: results}, nil
}
