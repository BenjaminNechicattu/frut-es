package devhandler

import (
	handlers "frutes/handler"

	"github.com/gofiber/fiber/v2"
)

func (dh *DevHandler) health(c *fiber.Ctx) error {

	response, err := dh.devSvc.HealthCheck()
	if err != nil {
		dh.logger.Error.Print(err)
		return handlers.APIResponseInternalServerError(c, "INTERNAL_SERVER_ERROR", "health check failed", err)
	}

	return handlers.APIResponseOK(c, response, "healthcheck finished")
}
