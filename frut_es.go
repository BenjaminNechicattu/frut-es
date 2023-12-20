package main

import (
	"frutes/config"
	mysqlservice "frutes/database/mysql"
	devhandler "frutes/handler/devapi"
	"frutes/server"
	devservice "frutes/services/dev"
	"frutes/utils"
	"log"
)

var (
	serverConfig *config.ServerConfig
	MysqlSvc     *mysqlservice.MysqlDBService
	logger       *utils.Logger
)

func init() {

	log.Println("initiating frut-es server")

	// parse conf
	log.Println("Parsing Configurations")

	var err error

	serverConfig, err = config.GetServerConfig()
	if err != nil {
		log.Panicln("EROR : ", err)
	}

	// initiate logger
	logger = utils.InitiateLoggers(serverConfig.LogFilePath)

	// initiate db
	MysqlSvc = mysqlservice.NewMysqlDBService(&serverConfig.MysqlConfig)
}

func main() {

	logger.Info.Println("starting server")

	// initiate services and their handlers
	devSvc := devservice.NewDevService(MysqlSvc, logger)
	devHandler := devhandler.NewDevHandler(devSvc, logger)

	// initiate server

	fiberApp := server.NewServer()

	handlers := []*server.ServerHandlerMap{
		server.NewServerHandlerMap("/dev/v1", devHandler),
	}

	server := &server.Server{
		Port:        *serverConfig.Port,
		FiberApp:    fiberApp,
		APIRootPath: "/api",
		Handlers:    handlers,
	}

	waitforShutdownInterrupt := server.StartServer()
	logger.Info.Println("Server started")

	select {
	case <-waitforShutdownInterrupt:
		goto stop
	}

stop:

	logger.Info.Println("initiating server stop sequence")

	server.ShutdownGracefully()

	logger.Info.Println("Server Stopped")
}
