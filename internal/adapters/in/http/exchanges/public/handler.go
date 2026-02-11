package public

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type exchangeService interface {
	Insert(ctx context.Context, comment string) error
	Get(ctx context.Context) ([]string, error)
}

type Handler struct {
	exchange exchangeService
}

func NewHandler(exchange exchangeService) *Handler {
	return &Handler{
		exchange: exchange,
	}
}

func (h *Handler) RegisterFastHTTPRoutes(r fiber.Router) {
	r.Post("/test", h.Insert)
	r.Get("/test", h.Get)
}

type testParams struct {
	Comment string `json:"comment"`
}

func (h *Handler) Insert(c *fiber.Ctx) error {
	var req testParams
	if err := c.ReqHeaderParser(&req); err != nil {
		return errors.Wrap(ErrInvalidRequestParams, err.Error())
	}
	if err := c.QueryParser(&req); err != nil {
		return errors.Wrap(ErrInvalidRequestParams, err.Error())
	}

	err := h.exchange.Insert(c.UserContext(), req.Comment)
	if err != nil {
		return errors.Wrap(err, "insert test")
	}

	return nil
}

func (h *Handler) Get(c *fiber.Ctx) error {
	tests, err := h.exchange.Get(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "get test")
	}

	return c.JSON(tests)
}
