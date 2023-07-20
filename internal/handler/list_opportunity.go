package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpportunityHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})

}
