package admin

import (
	"deto_api/internals/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup){	

	g.Use(middleware.AdminMiddleware())
	{
		
	}
	
}

