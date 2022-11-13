package link

import (
	"net/http"

	"../../domain"
	"../../services"
	"../../utils"
	"github.com/gin-gonic/gin"
)

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

func UrlClicked(c *gin.Context) {

	linkLog, restErr := services.UrlClicked()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, linkLog)
}
