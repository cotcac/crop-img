package router

import (
	"../img"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	router := r.Group("/api")
	{

		// CATEGORY ENDPOINT.

		router.POST("/img/", img.Insert)

	}
	return r
}

// Router is...
func Router() *gin.Engine {
	return setupRouter()
}
