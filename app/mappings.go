package app

import (
	"../controllers/ping"
	"../controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/user/find", users.FindUser)
	router.GET("/user/update", users.UpdateUser)
	router.POST("/user/create", users.CreateUser)
	router.POST("/linkData/create", users.CreateLinkData)
	// router.GET("/click", users.UrlClicked)
}
