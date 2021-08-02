package server

import (
	"fmt"
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func getPort() string {
	return fmt.Sprintf(":%s", os.Args[1])
}

var config = fiber.Config{
	DisableStartupMessage: true,
}

// Server ...
type Server struct {
	App      *fiber.App
	Listener net.Listener
}

// New ...
func New() Server {
	listener, err := net.Listen("tcp4", getPort())
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New(config)

	if os.Getenv("NODE_ENV") != "production" {
		app.Use(logger.New())
	}

	return Server{app, listener}
}

func (s *Server) Start() {
	s.App.Listener(s.Listener)
}
