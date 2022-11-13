package app

import (
	"../controllers/link"
	"../controllers/ping"
	"../controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/user/find", users.FindUser) // just CRUD operations
	router.GET("/user/update", users.UpdateUser)
	router.POST("/user/create", users.CreateUser)
	router.POST("/linkData/create", link.CreateLinkData) // API to create short URL
	router.POST("/cron/click", link.UrlClicked)          // API to update and store click time on DB
	router.GET("link/redirect", link.Redirect)           // API to redirect to a particular URL
}
