package model

import (
	"context"
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       uint32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"types:varchar(100);uniqueIndex;size:256"`
	UserId       uint32      `gorm:"types:int(11)"`
	UserCurrency string      `gorm:"types:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "orders"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
