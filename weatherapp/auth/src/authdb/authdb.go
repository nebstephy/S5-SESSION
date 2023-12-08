package authdb

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
}

func Connect(dbRoot string, dbPassword string, dbHost string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/", dbRoot, dbPassword, dbHost))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDB(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS auth")
	if err != nil {
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS auth.users (user_id int AUTO_INCREMENT, user_name char(50) NOT NULL, user_password char(128), PRIMARY KEY(user_id));")
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(db *sql.DB, user User) error {
	password := md5.Sum([]byte(user.Password))
	_, err := db.Exec("INSERT INTO auth.users (user_name,user_password) VALUES (?, ?)", user.Name, hex.EncodeToString(password[:]))
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(userName string, db *sql.DB) (User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM auth.users WHERE user_name = ?", userName)
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, u User) (bool, error) {
	user, err := GetUserByName(u.Name, db)
	if err != nil {
		return false, err
	}
	if user != (User{}) {
		return false, nil
	}

	err = InsertUser(db, u)
	if err != nil {
		return false, err
	}

	return true, nil
}
