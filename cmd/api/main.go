package main

import (
	"log"
	"net"

	"github.com/gofiber/fiber/v3"
)

func main() {
	host := getEnvironmentVariable(hostEnvironmentVariable, defaultHost)
	port := getEnvironmentVariable(portEnvironmentVariable, defaultPort)
	listenAddress := net.JoinHostPort(host, port)

	app := fiber.New()
	app.Get("/health", handleHealth)

	if listenError := app.Listen(listenAddress); listenError != nil {
		log.Fatalf("server stopped: %v", listenError)
	}
}

func handleHealth(requestContext fiber.Ctx) error {
	return requestContext.JSON(healthResponse{Status: "ok"})
}
