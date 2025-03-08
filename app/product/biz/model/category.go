package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	if name == "all" {
		err = c.db.WithContext(c.ctx).Model(&Category{}).Preload("Products").Find(&categories).Error
		return
	}
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}

func (c CategoryQuery) ListCategories() (categories []string, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Select("Name").Find(&categories).Error
	return
}

func (c CategoryQuery) AddCategory(name string) {
	err := c.db.WithContext(c.ctx).Create(&Category{Name: name}).Error
	if err != nil {
		panic(err)
	}
}

func (c CategoryQuery) DeleteCategory(name string) {
	err := c.db.WithContext(c.ctx).Delete(&Category{Name: name}).Error
	if err != nil {
		panic(err)
	}
}
