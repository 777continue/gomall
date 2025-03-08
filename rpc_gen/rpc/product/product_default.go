package product

import (
	"context"

	product "github.com/777continue/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = Client.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = Client.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = Client.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddProduct(ctx context.Context, req *product.AddProductReq, callOptions ...callopt.Option) (resp *product.AddProductResp, err error) {
	resp, err = Client.AddProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.UpdateProductReq, callOptions ...callopt.Option) (resp *product.UpdateProductResp, err error) {
	resp, err = Client.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *product.DeleteProductReq, callOptions ...callopt.Option) (resp *product.DeleteProductResp, err error) {
	resp, err = Client.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddCategory(ctx context.Context, req *product.AddCategoryReq, callOptions ...callopt.Option) (resp *product.AddCategoryResp, err error) {
	resp, err = Client.AddCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCategory(ctx context.Context, req *product.DeleteCategoryReq, callOptions ...callopt.Option) (resp *product.DeleteCategoryResp, err error) {
	resp, err = Client.DeleteCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListCategories(ctx context.Context, req *product.ListCategoriesReq, callOptions ...callopt.Option) (resp *product.ListCategoriesResp, err error) {
	resp, err = Client.ListCategories(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListCategories call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
