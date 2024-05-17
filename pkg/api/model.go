package api

import (
	"fmt"

	"eulabs/pkg/entity"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
	Price       string  `json:"price"`
	Currency    string  `json:"currency"`
}

func NewProductFromEntity(p *entity.Product, currencyCode string) *Product {
	if p == nil {
		return nil
	}

	price, ok := p.Prices[currencyCode]
	if !ok {
		return nil
	}

	return &Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		Price:       price.String(),
		Currency:    currencyCode,
	}
}

type ProductMutation struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	ImageURL    *string         `json:"image_url"`
	Prices      map[string]uint `json:"prices"`
}

func (pm *ProductMutation) Entity() *entity.Product {
	prices := make(map[string]entity.Currency, len(pm.Prices))
	for currencyCode, value := range pm.Prices {
		prices[currencyCode] = entity.Currency(value)
		fmt.Println(value, prices[currencyCode])
	}

	return &entity.Product{
		Name:        pm.Name,
		Description: pm.Description,
		ImageURL:    pm.ImageURL,
		Prices:      prices,
	}
}
