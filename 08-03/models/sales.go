package models

import (
	"time"

	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model

	Sale_date     time.Time `json:"sale_date"`
	Sale_time     time.Time `json:"sale_time"`
	Customer_name string    `json:"customer_name"`
	Total         float64   `json:"total"`
}

type GetSales struct {
	gorm.Model

	Sale_date     string  `json:"sale_date"`
	Sale_time     string  `json:"sale_time"`
	Customer_name string  `json:"customer_name"`
	Total         float64 `json:"total"`
}
