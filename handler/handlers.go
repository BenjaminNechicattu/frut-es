package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type APIHandler interface {
	RegisterRoutes(fiber.Router)
}

var ()

type APIResponseErrJson struct {
	ErrCode string      `json:"errcode"`
	ErrData interface{} `json:"errdata"`
}

type EmptyData struct{}

type APIResponse struct {
	Err  interface{} `json:"err"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func APIResponseOK(c *fiber.Ctx, data interface{}, msg string) error {
	responseobj := &APIResponse{
		Err:  nil,
		Data: data,
		Msg:  msg,
	}
	return c.Status(fiber.StatusOK).JSON(responseobj)
}

func APIResponseBadRequest(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusBadRequest).JSON(responseobj)
}

func APIResponseUnauthorized(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusUnauthorized).JSON(responseobj)
}

func APIResponseForbidden(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusForbidden).JSON(responseobj)
}

func APIResponseConflict(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusConflict).JSON(responseobj)
}

func APIResponseGone(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusGone).JSON(responseobj)
}

func APIResponseUnprocessableEntity(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusUnprocessableEntity).JSON(responseobj)
}

func APIResponseNotAcceptable(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusNotAcceptable).JSON(responseobj)
}

func APIResponseInternalServerError(c *fiber.Ctx, errorcode string, msg string, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(fiber.StatusInternalServerError).JSON(responseobj)
}

func APIFailedInternalAPICall(c *fiber.Ctx, errorcode string, msg string, statusCode int, errData interface{}) error {
	responseobj := &APIResponse{
		Err: &APIResponseErrJson{
			ErrCode: errorcode,
			ErrData: errData,
		},
		Data: nil,
		Msg:  msg,
	}
	return c.Status(statusCode).JSON(responseobj)
}
