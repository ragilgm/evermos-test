package repository

import (
	"evermos_technical_test/config"
	"evermos_technical_test/pkg/domain"
	"github.com/mitchellh/mapstructure"
)

type ProductRepository struct {
	config config.DbConfig
}

func MewProductRepository(config config.DbConfig) *ProductRepository {
	return &ProductRepository{
		config: config,
	}
}

func (p *ProductRepository) GetProducts() (*[]domain.Product, error) {

	var product []domain.Product

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&product).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	var respon *[]domain.Product

	// decode object respon dari database
	err = mapstructure.Decode(product, &respon)

	if err != nil {
		return nil, nil
	}
	return respon, nil
}

func (p ProductRepository) GetProductById(id uint) (res *domain.Product, err error) {

	var product domain.Product

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&product, id).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(product, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (p ProductRepository) UpdateProduct(id uint, req *domain.Product) (res *domain.Product, err error) {
	var product domain.Product

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&product, id).Error
	if err2 != nil {
		return nil, err2
	}

	err3 := connect.Model(&product).UpdateColumns(&req).Error
	if err3 != nil {
		return nil, err3
	}
	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(product, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (p ProductRepository) CreateProduct(req *domain.Product) (res *domain.Product, err error) {

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	_ = connect.Create(&req).RowsAffected

	// decode object respon dari database
	err = mapstructure.Decode(req, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (p ProductRepository) DeleteProduct(id int64) (err error) {
	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return err
	}
	// query to databse
	_ = connect.Delete(&domain.Product{}, id).RowsAffected

	return
}
