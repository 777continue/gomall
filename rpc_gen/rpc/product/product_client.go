package product

import (
	"context"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"

	"github.com/777continue/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error)
	UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
	AddCategory(ctx context.Context, Req *product.AddCategoryReq, callOptions ...callopt.Option) (r *product.AddCategoryResp, err error)
	DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *product.DeleteCategoryResp, err error)
	ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.ListCategoriesResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) AddProduct(ctx context.Context, Req *product.AddProductReq, callOptions ...callopt.Option) (r *product.AddProductResp, err error) {
	return c.kitexClient.AddProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, Req *product.UpdateProductReq, callOptions ...callopt.Option) (r *product.UpdateProductResp, err error) {
	return c.kitexClient.UpdateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) AddCategory(ctx context.Context, Req *product.AddCategoryReq, callOptions ...callopt.Option) (r *product.AddCategoryResp, err error) {
	return c.kitexClient.AddCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteCategory(ctx context.Context, Req *product.DeleteCategoryReq, callOptions ...callopt.Option) (r *product.DeleteCategoryResp, err error) {
	return c.kitexClient.DeleteCategory(ctx, Req, callOptions...)
}

func (c *clientImpl) ListCategories(ctx context.Context, Req *product.ListCategoriesReq, callOptions ...callopt.Option) (r *product.ListCategoriesResp, err error) {
	return c.kitexClient.ListCategories(ctx, Req, callOptions...)
}
