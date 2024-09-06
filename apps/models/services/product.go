package services

import (
	"product-ms/apps/models/repositories"
	"product-ms/apps/views"
	"product-ms/libs/helpers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(req views.CreateProductRequest) (views.Product, error)
	UpdateProductByID(id string, data views.UpdateProductRequest) (views.Product, error)
	DeleteProductByID(id string) error
	GetProductByID(id string) (views.Product, error)
	GetProducts(page, itemsPerPage int) ([]views.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepo
}

func NewProductService(productRepo repositories.ProductRepo) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProduct(req views.CreateProductRequest) (views.Product, error) {
	productExists, err := s.productRepo.GetProductByName(req.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return views.Product{}, err
    }

	if productExists != nil {
		return views.Product{}, helpers.ErrProductAlreadyExists
	}
	
	timeNow := time.Now()
	data := views.Product{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Variety:     req.Variety,
		Rating:      req.Rating,
		Stock:       req.Stock,
		CreatedAt:   &timeNow,
	}

	err = s.productRepo.InsertProduct(data)
	if err != nil {
		return views.Product{}, err
	}

	return data, nil
}

func (s *productService) UpdateProductByID(id string, data views.UpdateProductRequest) (views.Product, error) {
	_, err := s.productRepo.GetProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.Product{}, helpers.ErrProductNotFound
		}

		return views.Product{}, err
	}

	err = s.productRepo.UpdateProduct(id, data)
	if err != nil {
		return views.Product{}, err
	}

	updatedProduct, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return views.Product{}, err
	}

	return *updatedProduct, nil
}

func (s *productService) DeleteProductByID(id string) error {
	_, err := s.productRepo.GetProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.ErrProductNotFound
		}

		return err
	}

	return s.productRepo.DeleteProduct(id)
}

func (s *productService) GetProductByID(id string) (views.Product, error) {
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.Product{}, helpers.ErrProductNotFound
		}

		return views.Product{}, err
	}

	return *product, nil
}

func (s *productService) GetProducts(page, itemsPerPage int) ([]views.Product, error) {
	limit := itemsPerPage
	offset := (page - 1) * itemsPerPage

	products, err := s.productRepo.GetProducts(limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}
