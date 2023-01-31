package main

import (
	"fmt"
	"os"

	"github.com/aki-eiger/monorepo/app1/handlers"

	"github.com/gin-gonic/gin"
)

const azureFunctionsHandlerPort = "FUNCTIONS_CUSTOMHANDLER_PORT"

func main() {
	fmt.Println("Start!")
	g := gin.New()
	g.GET("/api/echo", handlers.Echo)

	port := resolvePort()
	err := g.Run(port)
	fmt.Println("Exited:", err)

}

func resolvePort() string {
	port := ":8080"
	if val, ok := os.LookupEnv(azureFunctionsHandlerPort); ok {
		port = ":" + val
	}
	return port
}
