// Code generated by hertz generator. DO NOT EDIT.

package MNproduct

import (
	MNproduct "github.com/777continue/gomall/app/frontend/biz/handler/MNproduct"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/categories", append(_listcategoriesMw(), MNproduct.ListCategories)...)
	_categories := root.Group("/categories", _categoriesMw()...)
	_categories.DELETE("/{category}", append(_deletecategoryMw(), MNproduct.DeleteCategory)...)
	root.POST("/categories", append(_addcategoryMw(), MNproduct.AddCategory)...)
	root.POST("/products", append(_addproductMw(), MNproduct.AddProduct)...)
	root.GET("/products", append(_listproductsMw(), MNproduct.ListProducts)...)
	_products := root.Group("/products", _productsMw()...)
	_products.PUT("/{id}", append(_updateproductMw(), MNproduct.UpdateProduct)...)
	_products.DELETE("/{id}", append(_deleteproductMw(), MNproduct.DeleteProduct)...)
}
