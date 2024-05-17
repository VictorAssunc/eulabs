package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nleof/goyesql"

	"eulabs/pkg/entity"
)

//go:embed queries/product.sql
var queriesFile []byte

type Product struct {
	db      *sql.DB
	queries goyesql.Queries
}

func NewProduct(db *sql.DB) *Product {
	return &Product{
		db:      db,
		queries: goyesql.MustParseBytes(queriesFile),
	}
}

func (r *Product) Create(ctx context.Context, product *entity.Product) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, r.queries["create-product"], product.Name, product.Description, product.ImageURL)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Println(txErr)
		}

		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Println(txErr)
		}

		return err
	}

	for currencyCode, value := range product.Prices {
		_, err = tx.ExecContext(ctx, r.queries["create-price"], id, value, currencyCode)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				log.Println(txErr)
			}

			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	product.ID = id
	return nil
}

func (r *Product) Get(ctx context.Context, id int64) (*entity.Product, error) {
	product := entity.NewEmptyProduct()
	row := r.db.QueryRowContext(ctx, r.queries["get-product"], id)
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.ImageURL); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, r.queries["get-price"], id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var currencyCode string
	var value entity.Currency
	for rows.Next() {
		if err = rows.Scan(&value, &currencyCode); err != nil {
			return nil, err
		}

		product.Prices[currencyCode] = value
	}

	return product, nil
}

func (r *Product) Update(ctx context.Context, product *entity.Product, removePrices map[string]struct{}) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.queries["update-product"], product.Name, product.Description, product.ImageURL, product.ID)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Println(txErr)
		}

		return err
	}

	for currencyCode := range removePrices {
		_, err = tx.ExecContext(ctx, r.queries["delete-price-by-currency"], product.ID, currencyCode)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				log.Println(txErr)
			}

			return err
		}
	}

	for currencyCode, value := range product.Prices {
		_, err = tx.ExecContext(ctx, r.queries["create-price"], product.ID, value, currencyCode)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				log.Println(txErr)
			}

			return err
		}
	}

	return tx.Commit()
}

func (r *Product) Delete(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, r.queries["delete-price"], id)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Println(txErr)
		}

		return err
	}

	_, err = tx.ExecContext(ctx, r.queries["delete-product"], id)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Println(txErr)
		}

		return err
	}

	return tx.Commit()
}
