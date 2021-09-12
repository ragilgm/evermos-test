package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	product2 "evermos_technical_test/pkg/dto/product"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
	"sync"
)

type ProductUseCase struct {
	mutex       sync.Mutex
	productRepo *repository.ProductRepository
}

func NewProductUseCase(mutex sync.Mutex, a *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		mutex:       mutex,
		productRepo: a,
	}
}

func (p ProductUseCase) GetProducts() (res *[]product2.ProductResponse, err error) {
	result, err := p.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(result, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p ProductUseCase) GetProductById(id uint) (res *product2.ProductResponse, err error) {
	result, err := p.productRepo.GetProductById(id)
	if err != nil {
		return nil, error2.ErrNotFound
	}

	err = mapstructure.Decode(result, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p ProductUseCase) UpdateProduct(id uint, req *product2.ProductRequest) (res *product2.ProductResponse, err error) {
	var product = domain.Product{}

	err = mapstructure.Decode(req, &product)
	if err != nil {
		return nil, err
	}

	result, err := p.productRepo.UpdateProduct(id, &product)
	if err != nil {
		return nil, error2.ErrNotFound
	}

	err = mapstructure.Decode(result, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p ProductUseCase) CreateProduct(req *product2.ProductRequest) (res *product2.ProductResponse, err error) {
	var product = domain.Product{}

	err = mapstructure.Decode(req, &product)
	if err != nil {
		return nil, err
	}

	result, err := p.productRepo.CreateProduct(&product)
	if err != nil {
		return nil, error2.ErrBadRequest
	}

	err = mapstructure.Decode(result, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p ProductUseCase) DeleteProduct(id int64) error {
	err := p.productRepo.DeleteProduct(id)
	if err != nil {
		return error2.ErrNotFound
	}
	return nil
}
