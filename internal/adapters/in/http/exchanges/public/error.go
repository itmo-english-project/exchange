package public

import "github.com/gofiber/fiber/v2"

var (
	ErrInvalidRequestParams = fiber.NewError(fiber.StatusBadRequest, "invalid request params")
)
