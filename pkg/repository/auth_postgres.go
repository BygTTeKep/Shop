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
	/*

		ПОДУМАТЬ КАК РЕАЛИЗОВАТЬ СОЗДАНИЕ КОРЗИНЫ ВО ВРЕМЯ СОЗДАНИЯ ПОЛЬЗОВАТЕЛЯ

	*/
	// tx, err := r.db.Begin()
	// if err != nil {
	// 	return -1, err
	// }
	// var idUser int
	// query := fmt.Sprintf("INSERT INTO %s(username, first_name, last_name, surname, email, password_hash) values($1, $2, $3, $4, $5, $6) RETURNING id", userTable)
	// row := tx.QueryRow(query, user.UserName, user.FirstName, user.LastName, user.Surname, user.Email, user.Password)
	// err = row.Scan(&idUser)
	// if err != nil {
	// 	tx.Rollback()
	// 	return -1, err
	// }
	// queryCart := fmt.Sprintf("INSERT INTO %s(customer_id) values $1 RETURNING id", cartTable)
	// defer tx.QueryRow(queryCart, idUser)
	// if err := rowCart.Scan(&idCart); err != nil {
	// 	tx.Rollback()
	// 	return -1, err
	// }
	// return idUser, tx.Commit()
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) DeleteUser(username, password string) error {
	user, err := r.GetUser(username, password)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err = r.db.Exec(query, user.Id)
	// if err := row.Scan(&user.Id); err != nil {
	// 	return -1, err
	// }
	return err
}
