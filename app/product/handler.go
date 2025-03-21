package main

import (
	"context"
	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	"github.com/777continue/gomall/app/product/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// AddProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) AddProduct(ctx context.Context, req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	resp, err = service.NewAddProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// AddCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) AddCategory(ctx context.Context, req *product.AddCategoryReq) (resp *product.AddCategoryResp, err error) {
	resp, err = service.NewAddCategoryService(ctx).Run(req)

	return resp, err
}

// DeleteCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	resp, err = service.NewDeleteCategoryService(ctx).Run(req)

	return resp, err
}

// ListCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListCategories(ctx context.Context, req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	resp, err = service.NewListCategoriesService(ctx).Run(req)

	return resp, err
}
