package router

import (
	"github.com/dyhalmeida/golang-opportunities/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handler.ListOpportunityHandler)
		v1.GET("/opening/:id", handler.ShowOpportunityHandler)
		v1.POST("/opening", handler.CreateOpportunityHandler)
		v1.PUT("/opening/:id", handler.UpdateOpportunityHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpportunityHandler)
	}
}
