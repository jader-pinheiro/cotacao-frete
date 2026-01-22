package v1

import (
	"log/slog"
	"strconv"

	"cotacao-fretes/internal/core/quote"
	"cotacao-fretes/internal/infra/requesthttp"
	"cotacao-fretes/internal/pkg/dto/requests"
	"cotacao-fretes/internal/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

type QuoteController struct {
	svc  *quote.Service
	clt  *requesthttp.Client
	slog *slog.Logger
}

func NewQuoteController(svc *quote.Service, clts *requesthttp.Client, slog *slog.Logger) *QuoteController {
	return &QuoteController{
		svc,
		clts,
		slog,
	}
}

func (ca *QuoteController) Get(ctx *fiber.Ctx) error {
	quote, err := ca.svc.Get(ctx.UserContext(), 1)

	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"msg":    err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(quote)

}

// Realiza a cotação de frete com base no payload enviado e salva a cotação no banco de dados
// @Summary Cotação de Frete
// @Tags Cotações
// @Param request body requests.RequestQuote  true "Payload com as informações necessárias para realizar a cotação de frete"
// @Success 201 {object} domain.Quote "Resposta com sucesso"
// @Failure 400 {object} errorResponse "Erro de requisição inválida"
// @Router /v1/quote [post]
func (ca *QuoteController) Insert(ctx *fiber.Ctx) error {
	var quotePayload requests.RequestQuote
	if err := ctx.BodyParser(&quotePayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"msg":    "invalid request payload",
		})
	}
	validation.ValidateQuoteRequest(quotePayload)

	if msgError := validation.ValidateQuoteRequest(quotePayload); msgError != "" {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"Msg":    msgError,
		})
	}
	req, err := ca.clt.GetQuoteWithPayload(quotePayload)

	quote, err := ca.svc.InsertQuote(ctx.UserContext(), req)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"msg":    err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(quote)
}

// Realiza a busca de cotações armazenadas com o resumo das cotações
// @Summary Resumo Cotação de Frete
// @Tags Cotações
// @Description Retorna um resumo das cotações, podendo limitar pelo número de últimas cotações
// @Param last_quotes query int false "Número de últimas cotações a serem retornadas"
// @Success 200 {array} domain.Quote "Lista de cotações"
// @Failure 400 {object} errorResponse "Erro de requisição inválida"
// @Router /v1/quote/metrics [get]
func (ca *QuoteController) GetResumeQuotes(ctx *fiber.Ctx) error {
	lastQuotesStr := ctx.Query("last_quotes")

	var lastQuotesInt *int

	if lastQuotesStr != "" {
		val, err := strconv.Atoi(lastQuotesStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "last_quotes must be a number",
			})
		}

		if val > 0 {
			lastQuotesInt = &val
		}
	}

	quote, err := ca.svc.GetResumeQuote(ctx.UserContext(), lastQuotesInt)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"msg":    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(quote)
}

type errorResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
