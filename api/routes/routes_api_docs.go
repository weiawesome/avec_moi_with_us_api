package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAPIDocsRoutes(r *gin.RouterGroup) {
	r.StaticFile("/api_docs.yaml", "templates/api_docs.yaml")
	r.GET("", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/redoc", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.HTML(http.StatusOK, "redoc-static.html", gin.H{})
	})
}
