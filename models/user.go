
package models

import (
	"fmt"
	"errors"
	"example.com/rest-api/db"
	"example.com/rest-api/utlities"
)

type  User struct{

	ID int64
	Email string `binding:"required"`
	Password   string `binding:"required"`
}

func (u User) Save() error{

	query := `INSERT INTO  users (email , password) VALUES (?,?)`
	stmt , err := db.DB.Prepare(query)

	if err!=nil {
		return err
	}

	defer stmt.Close()

	hashPassword , err := utlities.HashPassword(u.Password)
	if err!= nil {
		return err
	}
	result , err := stmt.Exec(u.Email, hashPassword)
	if err!= nil {
			return err
	}

	userId , err := result.LastInsertId()

	u.ID = userId
	return err


}

func (u *User) ValidateCredentials() error{

	query := "SELECT id,password FROM users WHERE email= ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID,&retrievePassword)

	if err != nil {
		fmt.Println(err)
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utlities.CheckPasswordHash(u.Password,retrievePassword)

	if !passwordIsValid {
		fmt.Println(err)
		return errors.New("Credentails invalid")
	}

	return nil


}