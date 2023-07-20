package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpportunityHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT opening by id: " + id,
	})

}
