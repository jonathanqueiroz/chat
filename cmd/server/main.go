package main

import (
	"log"
	"net/http"

	"github.com/jonathanqueiroz/chat/internal/app/routes"
	config "github.com/jonathanqueiroz/chat/internal/configs"
	"github.com/jonathanqueiroz/chat/internal/database"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Load configurations
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	db, err := database.NewDB(configs)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Routes
	routes.Init(e, db)

	// Start server
	e.Start(":8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
