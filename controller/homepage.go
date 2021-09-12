package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", radio)
	r.GET("/managesite", login)
	r.GET("/managesite/login", login)
	r.GET("/managesite/dashboard", dashboard)
	r.GET("/managesite/settings", settings)
	r.GET("/managesite/users", users)
	r.GET("/managesite/schedule", schedule)
}

func login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/login.tmpl",
		gin.H{
			"title": "Login",
		},
	)
}

func radio(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/radio.tmpl",
		gin.H{
			"title": "Radio Broadcast Al-Muwasholah",
		},
	)
}

func dashboard(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/dashboard.tmpl",
		gin.H{
			"title": "Dashboard",	
		},
	)
}

func settings(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/settings.tmpl",
		gin.H{
			"title": "Settings",
		},
	)
}

func users(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/users.tmpl",
		gin.H{
			"title": "Users",
		},
	)
}

func schedule(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/schedule.tmpl",
		gin.H{
			"title": "Schedule",
		},
	)
}
