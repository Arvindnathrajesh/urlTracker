package app

import (
	"fmt"
	"os"
	"urlTracker/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApp() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error in loading .env ")
	}
	MapUrls()
	domain.ConnDB()
	fmt.Println(os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
