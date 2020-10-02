package route

import (
	"github.com/gin-gonic/gin"
	"goshop-api/app/middleware"
)
import "goshop-api/app/controller"

func ServeRoutes() {
	router := gin.Default()
	users := router.Group("/api/users")
	{
		users.Use(middleware.Auth)
		users.GET("/", controller.GetUsers)
		users.GET("/detail/:id", controller.GetUser)
		users.DELETE("/:id", controller.DeleteUser)
		users.PUT("/:id", controller.UpdateUser)
		users.GET("/profile", controller.GetProfile)
	}

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}
	//run gin
	_ = router.Run()
}
