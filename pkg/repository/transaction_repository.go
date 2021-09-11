package repository

import (
	"evermos_technical_test/config"
	"evermos_technical_test/pkg/domain"
	"github.com/mitchellh/mapstructure"
)

type TransactionRepository struct {
	config config.DbConfig
}

func NewTransactionRepository(config config.DbConfig) *TransactionRepository {
	return &TransactionRepository{
		config: config,
	}
}

func (o TransactionRepository) CreateTransaction(req *domain.Transaction) (res *domain.Transaction, err error) {

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

func (o TransactionRepository) UpdateTransaction(id uint, req *domain.Transaction) (res *domain.Transaction, err error) {
	var transaction domain.Transaction
	// connect to databse
	connect, err := o.config.Connect()
	if err != nil {
		return nil, err
	}

	// connect to databse
	connect, err = o.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&transaction, id).Error
	if err2 != nil {
		return nil, err2
	}

	err3 := connect.Model(&transaction).Update(req).Error
	if err3 != nil {
		return nil, err3
	}
	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(req, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (p TransactionRepository) GetTransactionByOrderId(id uint) (res *domain.Transaction, err error) {
	var transaction domain.Transaction

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&transaction, "order_id=?", id).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(transaction, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}

func (p TransactionRepository) GetTransactionByToken(token string) (res *domain.Transaction, err error) {
	var transaction domain.Transaction

	// connect to databse
	connect, err := p.config.Connect()
	if err != nil {
		return nil, err
	}

	// query to databse
	err2 := connect.Find(&transaction, "transaction_token=?", token).Error
	if err2 != nil {
		return nil, err2
	}

	// close konneksi database
	connect.Close()

	// decode object respon dari database
	err = mapstructure.Decode(transaction, &res)

	if err != nil {
		return nil, nil
	}
	return res, nil
}
