package repository

import (
	"fmt"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, first_name, last_name, surname, email, password_hash) values($1, $2, $3, $4, $5, $6) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.UserName, user.FirstName, user.LastName, user.Surname, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := r.db.Exec(query, id)
	return err
}
