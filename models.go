package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
}
