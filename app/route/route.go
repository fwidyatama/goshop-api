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

	stores := router.Group("/api/stores")
	{
		stores.GET("/", controller.GetStores)
		stores.GET("/:id", controller.GetStore)
		stores.POST("/", controller.AddStore)
		stores.DELETE("/:id", controller.DeleteStore)
		stores.PUT("/:id", controller.UpdateStore)
		stores.POST("/additem", controller.AddStoreItem)
	}

	items := router.Group("/api/items")
	{
		items.POST("/", controller.AddItem)
		items.GET("/", controller.GetItems)
		items.GET("/:id", controller.GetItem)
		items.DELETE("/:id", controller.DeleteItem)
		items.PUT("/:id", controller.UpdateItem)
	}

	orders := router.Group("/api/order")
	{
		orders.Use(middleware.Auth)
		orders.POST("/", controller.AddOrder)
		orders.GET("/", controller.GetOrders)
	}

	//run gin
	_ = router.Run()
}
