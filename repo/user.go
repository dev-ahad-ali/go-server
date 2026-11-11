package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id
	`
	// execute named query
	var userID int
	rows, err := r.dbCon.NamedQuery(query, user)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID
	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User

	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`
	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
