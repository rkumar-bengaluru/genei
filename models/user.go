package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/logger"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Role     string
}

func (u User) Save(context *gin.Context) error {
	log := logger.Get(context).With(
		zap.String("username", u.Email),
		zap.String("role", u.Role),
		zap.String("method", "Save"))

	query := "INSERT INTO users(email, password,role) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.Role)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := result.LastInsertId()
	log.Info("User created with Id ", zap.Int64("id", userId))
	return err
}

func (u *User) ValidateCredentials(context *gin.Context) (string, error) {
	log := logger.Get(context).With(zap.String("username", u.Email), zap.String("role", u.Role))

	query := "SELECT id, password,role FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword, &u.Role)

	if err != nil {
		log.Error(err.Error())
		return "", errors.New(err.Error())
	}

	log.Info("reading user from db successful")
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return "", errors.New("credentials invalid")
	}

	return u.Role, nil
}
