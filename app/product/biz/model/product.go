package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Stock       uint32     `json:"stock"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error
	return
}

func (p ProductQuery) AddProduct(product *Product) {
	p.db.WithContext(p.ctx).Model(&Product{}).Create(product)
}

func (p ProductQuery) UpdateProduct(productId uint, updates map[string]interface{}) error {
	if stock, ok := updates["stock"]; ok {
		// 如果stock是负数，则减去该值
		if stock.(int32) < 0 {
			return p.db.WithContext(p.ctx).Model(&Product{}).
				Where("id = ?", productId).
				UpdateColumn("stock", gorm.Expr("stock - ?", -stock.(int32))).Error
		}
		// 如果stock是正数，则加上该值
		return p.db.WithContext(p.ctx).Model(&Product{}).
			Where("id = ?", productId).
			UpdateColumn("stock", gorm.Expr("stock + ?", stock.(int32))).Error
	}

	// 其他字段正常更新
	return p.db.WithContext(p.ctx).Model(&Product{}).
		Where("id = ?", productId).
		Select("*").
		Updates(updates).Error
}

func (p ProductQuery) DeleteProduct(productId int) {
	p.db.WithContext(p.ctx).Model(&Product{}).Delete(&Product{}, productId)
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func NewCachedProductQuery(pq ProductQuery, cacheClient *redis.Client) CachedProductQuery {
	return CachedProductQuery{productQuery: pq, cacheClient: cacheClient, prefix: "YZZ"}
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	// get from cache
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	// get from db
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, time.Hour)
	}
	return
}
