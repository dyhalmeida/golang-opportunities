package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpportunityHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening by id: " + id,
	})

}
