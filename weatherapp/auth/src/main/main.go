package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/abohmeed/auth/authdb"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const secretKey string = "xco0sr0fh4e52x03g9mv"

var dbHost string
var dbRoot = "root"
var dbPassword string

func main() {
	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}

	db := authdb.Connect(dbRoot, dbPassword, dbHost)
	if db == nil {
		fmt.Println("Could not connect to the database")
		return
	}
	defer db.Close()

	if err := authdb.CreateDB(db); err != nil {
		fmt.Println("Error creating database:", err)
		return
	}
	if err := authdb.CreateTables(db); err != nil {
		fmt.Println("Error creating tables:", err)
		return
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", health)
	router.POST("/users/:id", loginUser)
	router.POST("/users", createUser)

	router.Run(":8080")
}

// ... [rest of the functions remain the same]
