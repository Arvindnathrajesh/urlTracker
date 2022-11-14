package users

import (
	"net/http"

	"urlTracker/domain"
	"urlTracker/services"
	"urlTracker/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func FindUser(c *gin.Context) {
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := utils.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.FindUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userEmail := c.Query("email")
	field := c.Query("field")
	value := c.Query("value")
	if userEmail == "" {
		restErr := utils.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if field == "" {
		restErr := utils.BadRequest("no field.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if value == "" {
		restErr := utils.BadRequest("no value.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.UpdateUser(userEmail, field, value)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
