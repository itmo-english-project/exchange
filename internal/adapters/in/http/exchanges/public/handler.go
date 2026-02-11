package public

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/itmo-english-project/exchange/internal/adapters/out/postgres/exchanges"
)

type exchangeService interface {
	InsertExchange(ctx context.Context, fromID string, toID string, status exchanges.ExchangeStatus, comment string) error
	GetExchanges(ctx context.Context) ([]string, error)
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
	r.Post("/test", h.InsertExchangeOperation)
	r.Get("/test", h.Get)
}

type testParams struct {
	FromId  string                   `json:"from_id"`
	ToId    string                   `json:"to_id"`
	Status  exchanges.ExchangeStatus `json:"exchange_status"`
	Comment string                   `json:"comment"`
}

// InsertExchangeOperation inserts exchange event
//
//	@Summary	Inserts exchange event
//	@Accept		json
//	@Produce	json
//	@Param		from_id	query		string	false	"from id"
//	@Param		to_id	query		string	false	"to id"
//	@Param		exchange_status	query		exchanges.ExchangeStatus	false	"status"
//	@Param		comment	query		string	false	"comment"
//	@Success	200		{object} string "ok"
//	@Failure	500		{object} string "failure"
//	@Router		/test [post]
func (h *Handler) InsertExchangeOperation(c *fiber.Ctx) error {
	var req testParams
	if err := c.ReqHeaderParser(&req); err != nil {
		return errors.Wrap(ErrInvalidRequestParams, err.Error())
	}
	if err := c.QueryParser(&req); err != nil {
		return errors.Wrap(ErrInvalidRequestParams, err.Error())
	}

	err := h.exchange.InsertExchange(c.UserContext(), req.FromId, req.ToId, req.Status, req.Comment)
	if err != nil {
		return errors.Wrap(err, "insert test")
	}

	return nil
}

func (h *Handler) Get(c *fiber.Ctx) error {
	tests, err := h.exchange.GetExchanges(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "get test")
	}

	return c.JSON(tests)
}
