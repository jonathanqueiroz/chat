package routes

import (
	"github.com/jonathanqueiroz/tickets/internal/app/handlers"
	"github.com/jonathanqueiroz/tickets/internal/app/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	// Routes
	users := e.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
}
