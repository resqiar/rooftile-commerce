package main

import (
	"fmt"
	"log"
	"os"
	"resqiar/rooftile-commerce/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading .env file", err)
	}

	server := fiber.New()

	routers.InitWebRouter(server)

	PORT := os.Getenv("PORT")
	if err := server.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		log.Fatal(err)
	}
}
