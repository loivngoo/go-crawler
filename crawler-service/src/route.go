package main

import (
	"crawler-service/src/api"
	"crawler-service/src/component"

	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	v1Router(r, appCtx)
}

func v1Router(r *gin.Engine, appCtx component.AppContext) {
	// v1 := r.Group("/api/v1")
	v1 := r.Group("/")

	{
		v1.GET("/home", func(c *gin.Context) {
			c.JSON(200, "homepage")
		})

		// v1.GET("/lotto", func(c *gin.Context) {
		v1.GET("/", func(c *gin.Context) {
			defer api.StartCrawl()
			c.JSON(200, "lotto")
		})
	}

}
