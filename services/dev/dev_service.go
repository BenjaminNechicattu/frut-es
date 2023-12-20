package devservice

import (
	mysqlservice "frutes/database/mysql"
	"frutes/utils"
)

const (
	ORG_INVITE_EXPIRE_TIME = 48 // in hours
)

type DevService struct {
	mysqlDBService *mysqlservice.MysqlDBService
	logger         *utils.Logger
}

func NewDevService(mysqlDBService *mysqlservice.MysqlDBService, logger *utils.Logger) *DevService {
	return &DevService{
		mysqlDBService: mysqlDBService,
		logger:         logger,
	}
}
