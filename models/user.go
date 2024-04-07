package models

import (
	"errors"

	"github.com/isaacchunn/rest-api/db"
	"github.com/isaacchunn/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//Hash user password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	//Exec statement
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	//Unique emails
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials1")
	}
	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("invalid credentials2")
	}
	return nil
}
