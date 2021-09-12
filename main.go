package main

import (
	"log"
	"hel/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/managesite/css", "./static/css")
	r.Static("/managesite/images", "./static/images")
	r.Static("/managesite/scss", "./static/scss")
	r.Static("/managesite/vendors", "./static/vendors")
	r.Static("/managesite/fonts", "./static/fonts")
	r.Static("/managesite/js", "./static/js")
	r.StaticFile("/managesite/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("templates/**/*")
	controller.Router(r)

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0`:8080 (for windows "localhost:8080")
}
