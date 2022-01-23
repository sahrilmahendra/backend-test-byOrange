package routes

import (
	"byorange/constants"
	"byorange/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// function routes
func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// route users tanpa JWT
	e.POST("/signup", controllers.CreateUserControllers)
	e.POST("/login", controllers.LoginUserControllers)

	// route items tanpa JWT
	e.GET("/items", controllers.GetAllItemControllers)
	e.GET("/items/:id", controllers.GetItemByIdControllers)

	// route order
	e.POST("/order", controllers.CreatePurchaseOrderControllers)

	// group JWT
	j := e.Group("/jwt")
	j.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// route users dengan JWT
	j.GET("/users", controllers.GetAllUsersControllers)
	j.GET("/users/:id", controllers.GetUserControllers)
	j.PUT("/users/:id", controllers.UpdateUserControllers)
	j.DELETE("/users/:id", controllers.DeleteUserControllers)

	// route items dengan JWT
	j.POST("/items", controllers.CreateItemControllers)
	j.PUT("/items/:id", controllers.UpdateItemControllers)
	j.DELETE("/items/:id", controllers.DeleteItemControllers)

	return e
}
