package utils

import (
	"database/sql"
	"go-next-memo/database"
	"go-next-memo/models"
)

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	db := database.GetDB()
	row := db.QueryRow("SELECT email, password FROM user WHERE email = ? ", email)
	
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err // No user found
		}
		return user, err // Other errors
	}

	return user, nil
}

