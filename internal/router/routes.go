package router

import (
	docs "github.com/dyhalmeida/golang-opportunities/docs"
	"github.com/dyhalmeida/golang-opportunities/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ListOpportunityHandler)
		v1.GET("/opening/:id", handler.ShowOpportunityHandler)
		v1.POST("/opening", handler.CreateOpportunityHandler)
		v1.PUT("/opening/:id", handler.UpdateOpportunityHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpportunityHandler)
	}
	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
