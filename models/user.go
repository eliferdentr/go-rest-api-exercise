package models

import (
	"fmt"

	"eliferden.com/restapi/db"
	"eliferden.com/restapi/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
    Salt string
}

func (user User) Save() error{
	query := "INSERT INTO USERS (EMAIL, PASSWORD) VALUES ($1, $2) RETURNING ID"
    var id int64
    
    hashedPassword, _ , err := utils.HashPasswordArgon2id(user.Password)

    if err != nil {
        fmt.Println(err)
        return err
    }

    err = db.DB.QueryRow(query, user.Email, hashedPassword).Scan(&id)

    if err != nil {
        fmt.Println(err)
        return err
    }
    user.ID = id
    return err

}