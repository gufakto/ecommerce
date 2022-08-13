package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gufakto/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addProduct", controllers.AddProduct())
	incomingRoutes.GET("/users/productView", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/search", controllers.SearchProduct())
}
