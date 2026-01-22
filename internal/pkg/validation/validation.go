package validation

import (
	"cotacao-fretes/internal/pkg/dto/requests"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customMessages = map[string]string{
	"required": "O campo {field} é obrigatório.",
}

func replacePlaceholders(message, field, param string) string {
	message = strings.ReplaceAll(message, "{field}", field)
	message = strings.ReplaceAll(message, "{param}", param)
	return message
}

func checkStruct(e error) string {
	var messages []string
	var message string
	if validationErrors, ok := e.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()
			param := err.Param()
			customMessage, exists := customMessages[tag]
			if !exists {
				customMessage = "O campo {field} é inválido."
			}

			finalMessage := replacePlaceholders(customMessage, field, param)
			messages = append(messages, finalMessage)
		}

		message = strings.Join(messages, ", ")
	}

	return message
}

func ValidateQuoteRequest(quotePayload requests.RequestQuote) string {

	validate := validator.New()
	var message string

	if err := validate.Struct(&quotePayload); err != nil {
		message = checkStruct(err)
	}

	return message
}
