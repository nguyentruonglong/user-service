package main

import (
	"log"
	"os"

	"user_service/auth"
	db "user_service/db"
	routers "user_service/routers"
	users "user_service/versioned/v1/users"

	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	//Start SQLite3 database
	sqliteDB, err := db.Init()
	sqliteDB.AutoMigrate(&users.UserModel{})

	port := os.Getenv("PORT")

	r := gin.New()
	r.Use(auth.CORSMiddleware())
	r.Use(auth.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	routers.BindRouters(r)

	r.Run(":" + port)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {

		SSLKeys := &struct {
			CERT string
			KEY  string
		}{}

		// Generated using sh generate-certificate.sh
		SSLKeys.CERT = "./cert/ca.cer"
		SSLKeys.KEY = "./cert/ca.key"

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}

}
