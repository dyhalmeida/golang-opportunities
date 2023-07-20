package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpportunityHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening by id: " + id,
	})

}
