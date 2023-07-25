package handler

import (
	"net/http"

	"github.com/dyhalmeida/golang-opportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpportunityHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "routerParameter").Error())
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening with id: "+id+" not found")
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting opening with id: "+id)
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
