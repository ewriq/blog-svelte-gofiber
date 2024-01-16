package Database

import (
	"basic-blog/Utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Users(Token string) (User, error) {
	query := "SELECT * FROM user WHERE token = ?"
	row := db.QueryRow(query, Token)

	var user User
	err := row.Scan(&user.Username, &user.Password, &user.Email, &user.Token, &user.Perm)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("kullanıcı bulunamadı")
		}
		return User{}, err
	}
	return user, nil
}

func Login(email, password string) (string, error) {
	query := "SELECT token FROM user WHERE email = ? AND password = ?"
	row := db.QueryRow(query, email, password)

	var token string
	if err := row.Scan(&token); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Kullanıcı bulunamadı")
		}
		return "", err
	}

	return token, nil
}

func Register(email, password, username string) (bool, string) {
	Token := Utils.Token(10)
	if !FinderEmail(email) {
		query := "INSERT INTO user (username, password, email, token) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(query, username, password, email, Token)
		if err != nil {
			fmt.Println(err)
			return false, ""
		}
		return true,Token
	} else {
		return false,""
	}
}

func FinderEmail(email string) bool {
	query := "SELECT COUNT(*) FROM user WHERE email = ?"
	row := db.QueryRow(query, email)

	var user int
	err := row.Scan(&user)
	if err != nil {
		return false
	}

	if user == 0 {
		return false
	}

	return true
}
