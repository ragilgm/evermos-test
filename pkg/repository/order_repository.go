package repository

import (
	"evermos_technical_test/config"
	"evermos_technical_test/pkg/domain"
	"github.com/mitchellh/mapstructure"
)

type OrderRepository struct {
	config config.DbConfig
}

func NewOrderRepository(config config.DbConfig) *OrderRepository {
	return &OrderRepository{
		config: config,
	}
}

func (o OrderRepository) CreateOrder(req domain.Order) (res *domain.Order, err error) {

	// connect to databse
	connect, err := o.config.Connect()
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

func (o OrderRepository) UpdateOrder(id int64, req domain.Order) (res *domain.Order, err error) {
	var order domain.Order

	// connect to databse
	connect, err := o.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&order, id).Error
	if err2 != nil {
		return nil, err2
	}

	err3 := connect.Model(&order).Update(req).Error
	if err3 != nil {
		return nil, err3
	}
	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(order, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (o OrderRepository) GetOrderById(id uint) (res *domain.Order, err error) {

	var order domain.Order

	// connect to databse
	connect, err := o.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Preload("OrderItems").Preload("OrderItems.Item").Find(&order, id).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(order, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}
