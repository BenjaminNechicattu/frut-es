package devhandler

import (
	devsvc "frutes/services/dev"
	"frutes/utils"

	"github.com/gofiber/fiber/v2"
)

type DevHandler struct {
	devSvc *devsvc.DevService
	logger *utils.Logger
}

func (dh *DevHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/health", dh.health)
}

func NewDevHandler(devSvc *devsvc.DevService, logger *utils.Logger) *DevHandler {
	return &DevHandler{
		devSvc: devSvc,
		logger: logger,
	}
}
