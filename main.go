package main

import (
	"fmt"
	"log"
	"os"
	"resqiar/rooftile-commerce/configs"
	"resqiar/rooftile-commerce/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading .env file", err)
	}

	// initialize database connection
	pool := configs.InitDBConfig()
	defer pool.Close()

	// define where the views (frontend) files are
	engine := html.New("./views/pages", ".html")
	server := fiber.New(fiber.Config{
		Views: engine,
	})

	// initialize routers
	routers.InitWebRouter(server)

	PORT := os.Getenv("PORT")
	if err := server.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		log.Fatal(err)
	}
}
