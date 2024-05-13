package main

import (
	"database/sql"
	// "log"
	// "net/http"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "minpro-backend/models"
	"project/controller"
	"project/config"
)

var DB *sql.DB

func main() {
	e := echo.New()

	// Initialize MySQL connection
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	// Auto migrate schema
	err = config.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	// Routes
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.GET("/books", controller.GetBook)
	e.GET("/books/:ID", controller.GetBookById)
	e.POST("/books", controller.CreateBook)
	e.PUT("/books/:ID", controller.UpdateBook)
	e.DELETE("/books/:ID", controller.DeleteBook)

	// Start server
	e.Start(":8080")
}
