package handler

import (
	"net/http"

	"github.com/dyhalmeida/golang-opportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpportunityHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "routerParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening with id: "+id+" not found")
		return
	}
	sendSuccess(ctx, "show-opening", opening)
}
