package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-backend-template/internal/database"
	v1 "github.com/pjchender/go-backend-template/internal/router/api/v1"
)

func RegisterOrganization(db *database.GormDatabase, routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	orgHandler := v1.NewOrganizationHandler(db)
	orgRouter := routerGroup.Group("/organizations", middleware...)
	{
		orgRouter.GET("/:id", orgHandler.Get)
	}
}
