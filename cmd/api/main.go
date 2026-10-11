package main

import (
	"log"
	"net"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"

	_ "github.com/basiliok/matrix-qr-api-go/docs"
)

// @title       matrix-qr-api-go
// @version     1.0
// @description QR factorization of a matrix and its statistics.
func main() {
	host := getEnvironmentVariable(hostEnvironmentVariable, defaultHost)
	port := getEnvironmentVariable(portEnvironmentVariable, defaultPort)
	listenAddress := net.JoinHostPort(host, port)

	app := fiber.New()
	app.Get("/health", handleHealth)
	app.Get("/docs/*", swaggo.HandlerDefault)

	if listenError := app.Listen(listenAddress); listenError != nil {
		log.Fatalf("server stopped: %v", listenError)
	}
}

// @Summary  Check that the service is running
// @Tags     Health
// @Produce  json
// @Success  200 {object} healthResponse
// @Router   /health [get]
func handleHealth(requestContext fiber.Ctx) error {
	return requestContext.JSON(healthResponse{Status: "ok", Service: "matrix-qr-api-go"})
}
