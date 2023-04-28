package repository

import (
	"fmt"
	"strings"

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

func (r *AuthPostgres) UpdateUser(id int, user model.User) error {
	setvalues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if user.FirstName != "" {
		setvalues = append(setvalues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, user.FirstName)
		argId++
	}
	if user.LastName != "" {
		setvalues = append(setvalues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, user.LastName)
		argId++
	}
	if user.Surname != "" {
		setvalues = append(setvalues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, user.Surname)
		argId++
	}
	if user.Email != "" {
		setvalues = append(setvalues, fmt.Sprintf("email=$%d", argId))
		args = append(args, user.Email)
		argId++
	}
	if user.Password != "" {
		setvalues = append(setvalues, fmt.Sprintf("password_hash=$%d", argId))
		args = append(args, user.Password)
		argId++
	}
	setQuery := strings.Join(setvalues, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", userTable, setQuery, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}
