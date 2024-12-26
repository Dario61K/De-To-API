package routes

import (
	"deto_api/internals/admin"
	"deto_api/internals/auth"
	"deto_api/internals/business"
	"deto_api/internals/product"
	"deto_api/internals/rating"
	"deto_api/internals/search"
	"deto_api/internals/user"

	"github.com/gin-gonic/gin"
)

func SetupRutes() *gin.Engine {

	r := gin.Default()

	admin.RegisterRoutes(r.Group("/admin"))
	user.RegisterRoutes(r.Group("/user"))
	auth.RegisterRoutes(r.Group("/auth"))
	business.RegisterRoutes(r.Group("/business"))
	product.RegisterRoutes(r.Group("/product"))
	rating.RegisterRoutes(r.Group("/rating"))
	search.RegisterRoutes(r.Group("/search"))
	
	return r
}

