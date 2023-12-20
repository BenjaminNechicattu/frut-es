package server

import (
	"context"
	"encoding/json"
	handlers "frutes/handler"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type ServerHandlerMap struct {
	APIPath string
	Handler handlers.APIHandler
}

type Server struct {
	Port        int
	FiberApp    *fiber.App
	APIRootPath string
	Handlers    []*ServerHandlerMap
}

func NewServerHandlerMap(apipath string, handler handlers.APIHandler) *ServerHandlerMap {
	return &ServerHandlerMap{APIPath: apipath, Handler: handler}
}

func (s *Server) Setup() {
	if s.FiberApp == nil {
		log.Fatalln("Server setup incorrectly!")
	}
	AddMiddlewares(s.FiberApp)

	rootgroup := s.FiberApp.Group(s.APIRootPath)

	for _, eachhandlermap := range s.Handlers {
		apigroup := rootgroup.Group(eachhandlermap.APIPath)
		eachhandlermap.Handler.RegisterRoutes(apigroup)
	}
}

func (s *Server) StartServer() <-chan os.Signal {
	s.Setup()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.FiberApp.Listen(":" + strconv.Itoa(s.Port)); err != nil {
			log.Fatal(err)
		}
	}()
	return quit
}

func NewServer() *fiber.App {
	config := fiber.Config{
		ReadTimeout: 2 * time.Second,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
	return fiber.New(config)
}

func (s *Server) ShutdownGracefully() {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		// Release resources like Database connections
		cancel()
	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- s.FiberApp.Shutdown() }()

	select {
	case <-timeout.Done():
		log.Fatal("Server Shutdown Timed out before shutdown.")
	case err := <-shutdownChan:
		if err != nil {
			log.Fatal("Error while shutting down server", err)
		} else {
			log.Println("Server Shutdown Successful")
		}
	}
}

func AddMiddlewares(a *fiber.App) {
	a.Use(
		cors.New(),
		compress.New(),
		recover.New(),
		requestid.New(),
		logger.New(logger.Config{
			Format: "${time} ${pid} ${locals:requestid} ${status} - ${method} ${path} [${latency}] \n${reqHeaders} ${body}\n",
		}),
	)
}
