package users

import (
	"net/http"

	"../../domain"
	"../../services"
	"../../utils"
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

func CreateLinkData(c *gin.Context) {
	var newLinkData domain.LinkData
	if err := c.ShouldBindJSON(&newLinkData); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}

	linkData, restErr := services.CreateLinkData(&newLinkData)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, linkData)
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

// func UrlClicked(c *gin.Context) {
// 	shortUrl := c.Query("shortUrl")
// 	usrPhone := c.Query("phone")
// 	linkData, restErr1 := services.FindLinkData(shortUrl)
// 	if restErr1 != nil {
// 		c.JSON(restErr1.Status, restErr1)
// 		return
// 	}
// 	fmt.Println(linkData)
// 	userLinkData, restErr := services.UrlClicked(linkData.Url, usrPhone)
// 	if restErr != nil {
// 		c.JSON(restErr.Status, restErr)
// 		return
// 	}

// 	c.JSON(http.StatusOK, userLinkData)
// }

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
