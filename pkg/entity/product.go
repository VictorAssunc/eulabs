package entity

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	ImageURL    *string `db:"image_url"`

	Prices map[string]Currency
}

func NewEmptyProduct() *Product {
	return &Product{
		Prices: make(map[string]Currency),
	}
}
