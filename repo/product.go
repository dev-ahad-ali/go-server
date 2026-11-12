package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	return &productRepo{
		db: dbCon,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
			titile,
			description,
			price,
			img_url
		)
		VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		FROM products
		WHERE id = $1
	`
	err := r.db.Get(&prd, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	query := `
		SELECT
			id,
			title,
			description,
			price,
			img_url
		FROM products
	`
	err := r.db.Select(&prdList, query)

	if err != nil {
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Delete(productId int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID == productId {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {

	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}
