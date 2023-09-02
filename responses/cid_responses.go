package responses

import (
	"cid-10-api/models"

	"github.com/gofiber/fiber/v2"
)

type CidResponse struct {
	Status 		int 			`json:"status"`
	Message 	string			`json:"message"`
	Data 		models.Cid		`json:"data,omitempty"`
	Response	*fiber.Map 		`json:"response,omitempty"`
}