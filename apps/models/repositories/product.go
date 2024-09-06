package repositories

import (
	"product-ms/apps/views"

	"gorm.io/gorm"
)

type ProductRepo interface {
	InsertProduct(data views.Product) error
	UpdateProduct(id string, data views.UpdateProductRequest) error
	DeleteProduct(id string) error
	GetProductByID(id string) (*views.Product, error)
	GetProductByName(name string) (*views.Product, error)
	GetProducts(limit, offset int) ([]views.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) InsertProduct(data views.Product) error {
	return r.db.Create(&data).Error
}

func (r *productRepo) UpdateProduct(id string, data views.UpdateProductRequest) error {
	return r.db.Model(&views.Product{}).Where("id = ?", id).Updates(data).Error
}

func (r *productRepo) DeleteProduct(id string) error {
	return r.db.Delete(&views.Product{}, id).Error
}

func (r *productRepo) GetProductByID(id string) (*views.Product, error) {
	var product views.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) GetProductByName(name string) (*views.Product, error) {
	var product views.Product
	err := r.db.Where("name = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) GetProducts(limit, offset int) ([]views.Product, error) {
	var db *gorm.DB
	if limit == 0 || offset == 0 {
		db = r.db
	} else {
		db = r.db.Limit(limit).Offset(offset)
	}

	var products []views.Product
	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
