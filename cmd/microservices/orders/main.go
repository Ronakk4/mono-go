package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	  "github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	log.Println("Starting orders microservice")
  	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default environment variables")
    }
	
	bindAddr := os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR")

	if bindAddr == "" {
		bindAddr = ":3000" 
	}



	
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	
	go func() {
		if err := app.Listen(bindAddr); err != nil  {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	
	select {
	case sig := <-signalChan:
		log.Printf("Received signal: %v, shutting down...", sig)
	}


	if err := app.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}

 func createOrdersMicroservice(app *fiber.App) {
}
