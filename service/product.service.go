package service

import (
	"database/sql"
	"ecommerce/entity"
	"errors"
	"net/http"

	"github.com/lib/pq"
)

type ProductService interface {
	SaveProduct(entity.Product) (entity.Product, int, error)
	UpdateProduct(entity.Product, string) (entity.Product, int, error)
}

type productService struct {
	db       *sql.DB
	products []entity.Product
}

func NewProductService(db *sql.DB) ProductService {
	return &productService{
		db: db,
	}
}

func (service *productService) SaveProduct(product entity.Product) (entity.Product, int, error) {
	_, errdb := service.db.Exec(
		"insert into products(name,price,image_url,stock,condition,tags,is_purchasable) values ($1,$2,$3,$4,$5,$6,$7)",
		product.Name,
		product.Price,
		product.ImageUrl,
		product.Stock,
		product.Condition,
		pq.Array(product.Tags),
		product.IsPurchasable,
	)
	if errdb, ok := errdb.(*pq.Error); ok {
		return product, http.StatusInternalServerError, errdb
	}
	return product, http.StatusOK, nil
}

func (service *productService) UpdateProduct(product entity.Product, id string) (entity.Product, int, error) {
	res, errdb := service.db.Exec(
		"update products set name=$1, price=$2, image_url=$3, stock=$4, condition=$5, tags=$6, is_purchasable=$7 where product_id=$8",
		product.Name,
		product.Price,
		product.ImageUrl,
		product.Stock,
		product.Condition,
		pq.Array(product.Tags),
		product.IsPurchasable,
		id,
	)
	affRows, err := res.RowsAffected()
	if err != nil {
		return product, http.StatusInternalServerError, err
	}

	if affRows < 1 {
		return product, http.StatusNotFound, errors.New("product not found")
	}

	if errdb, ok := errdb.(*pq.Error); ok {
		return product, http.StatusInternalServerError, errdb
	}
	return product, http.StatusOK, nil
}
