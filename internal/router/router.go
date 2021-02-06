package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-backend-template/internal/database"
	"github.com/pjchender/go-backend-template/internal/middleware"
)

func New(db *database.GormDatabase) *gin.Engine {
	router := gin.New()
	router.RedirectTrailingSlash = false

	router.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler(), middleware.ResponseHeader())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		RegisterUser(db, v1)
		RegisterOrganization(db, v1)
	}

	return router
}
