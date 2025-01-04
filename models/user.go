package models

import (
	"errors"
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
	query := "INSERT INTO USERS (EMAIL, PASSWORD, SALT) VALUES ($1, $2, $3) RETURNING ID"
    var id int64
    
    hashedPassword, salt , err := utils.HashPasswordArgon2id(user.Password)

    if err != nil {
        fmt.Println(err)
        return err
    }

    err = db.DB.QueryRow(query, user.Email, hashedPassword, salt).Scan(&id)

    if err != nil {
        fmt.Println(err)
        return err
    }
    user.ID = id
    return err
}

func (u *User) ValidateCredentials() error{
    query := "SELECT PASSWORD, SALT, ID FROM USERS WHERE EMAIL = $1"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    var retrievedSalt string
    err := row.Scan(&retrievedPassword, &retrievedSalt, &u.ID)

    if err != nil {
        return err
    }

    isPasswordValid := utils.CompareHashArgon2id(u.Password, retrievedSalt, retrievedPassword)

    if !isPasswordValid {
        return errors.New("Credentials invalid")
    }
    return nil
}