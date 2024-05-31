package utils

import (
	"database/sql"
	"go-next-memo/database"
	"go-next-memo/models"
)

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	db := database.GetDB()
	row := db.QueryRow("SELECT email, password FROM user where email = ? ", user.Email)

	if row.Err() == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}