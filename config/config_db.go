package config

import (
	"evermos_technical_test/pkg/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbConfig struct {
	DB *gorm.DB
}

var err error

func (db *DbConfig) Connect() (*gorm.DB, error) {

	db.DB, err = gorm.Open("postgres", "host=localhost user=postgres password=Ragil404* dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	// migrate database
	db.DB.AutoMigrate(&domain.Product{})
	db.DB.AutoMigrate(&domain.Order{})
	db.DB.AutoMigrate(&domain.OrderItem{})
	db.DB.AutoMigrate(&domain.Item{})
	db.DB.AutoMigrate(&domain.Transaction{})
	return db.DB, nil
}
