package service

import (
	"context"
	"errors"
	"fmt"

	"eulabs/pkg/entity"
	"eulabs/pkg/repository"
)

var (
	ErrInvalidCurrencyCode = errors.New("invalid currency code")
)

type Product struct {
	repository *repository.Product
}

func NewProduct(repo *repository.Product) *Product {
	return &Product{
		repository: repo,
	}
}

func (s *Product) Create(ctx context.Context, product *entity.Product) error {
	for currencyCode := range product.Prices {
		if len(currencyCode) != 3 {
			return fmt.Errorf("%w: %s", ErrInvalidCurrencyCode, currencyCode)
		}
	}

	return s.repository.Create(ctx, product)
}

func (s *Product) Get(ctx context.Context, id int64) (*entity.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *Product) Update(ctx context.Context, oldProduct, product *entity.Product) error {
	for currencyCode := range product.Prices {
		if len(currencyCode) != 3 {
			return fmt.Errorf("%w: %s", ErrInvalidCurrencyCode, currencyCode)
		}
	}

	removePrices := make(map[string]struct{}, len(oldProduct.Prices))
	for currencyCode := range oldProduct.Prices {
		if _, ok := product.Prices[currencyCode]; !ok {
			removePrices[currencyCode] = struct{}{}
		}
	}

	return s.repository.Update(ctx, product, removePrices)

}

func (s *Product) Delete(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}
