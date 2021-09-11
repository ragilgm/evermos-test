package usecase

import (
	error2 "evermos_technical_test/error"
	"evermos_technical_test/pkg/domain"
	product2 "evermos_technical_test/pkg/dto/product"
	"evermos_technical_test/pkg/repository"
	"github.com/mitchellh/mapstructure"
)

type ProductUseCase struct {
	productRepo *repository.ProductRepository
}

func NewProductUseCase(a *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: a,
	}
}

func (p ProductUseCase) GetProducts() (res *[]product2.ProductResponse, err error) {
	result, err := p.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(result, &res)
	return res, nil
}

func (p ProductUseCase) GetProductById(id uint) (res *product2.ProductResponse, err error) {
	result, err := p.productRepo.GetProductById(id)
	if err != nil {
		return nil, error2.ErrNotFound
	}

	mapstructure.Decode(result, &res)
	return res, nil
}

func (p ProductUseCase) UpdateProduct(id uint, req *product2.ProductRequest) (res *product2.ProductResponse, err error) {
	var product = domain.Product{}

	mapstructure.Decode(req, &product)

	result, err := p.productRepo.UpdateProduct(id, &product)
	if err != nil {
		return nil, error2.ErrNotFound
	}

	mapstructure.Decode(result, &res)
	return res, nil
}

func (p ProductUseCase) CreateProduct(req *product2.ProductRequest) (res *product2.ProductResponse, err error) {
	var product = domain.Product{}

	mapstructure.Decode(req, &product)

	result, err := p.productRepo.CreateProduct(&product)
	if err != nil {
		return nil, error2.ErrBadParamInput
	}

	mapstructure.Decode(result, &res)

	return res, nil
}

func (p ProductUseCase) DeleteProduct(id int64) error {
	err := p.productRepo.DeleteProduct(id)
	if err != nil {
		return error2.ErrNotFound
	}
	return nil
}
