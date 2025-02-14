package v1

import (
	"github.com/gofiber/fiber/v3"

	utils_v1 "github.com/FDSAP-Git-Org/hephaestus/utils/v1"
)

func JSONResponse(c fiber.Ctx, retCode, retMessage string, httpStatusCode int) error {
	return c.Status(httpStatusCode).JSON(fiber.Map{
		"responseTime": utils_v1.GetResponseTime(c),
		"device":       string(c.RequestCtx().UserAgent()),
		"retCode":      retCode,
		"message":      retMessage,
	})
}

func JSONResponseWithData(c fiber.Ctx, retCode, retMessage string, data interface{}, httpStatusCode int) error {
	return c.Status(httpStatusCode).JSON(fiber.Map{
		"responseTime": utils_v1.GetResponseTime(c),
		"device":       string(c.RequestCtx().UserAgent()),
		"retCode":      retCode,
		"message":      retMessage,
		"data":         data,
	})
}

func JSONResponseWithError(c fiber.Ctx, retCode, retMessage string, err error, httpStatusCode int) error {
	return c.Status(httpStatusCode).JSON(fiber.Map{
		"responseTime": utils_v1.GetResponseTime(c),
		"device":       string(c.RequestCtx().UserAgent()),
		"retCode":      retCode,
		"message":      retMessage,
		"error":        err,
	})
}
