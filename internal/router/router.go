package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	addr := ":3333"
	r.Run(addr)
}
