package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"xor-operation/dto"
	"xor-operation/domain"
	"time"
	"net/http"
)

type EncryptAPI struct {
	encryptService domain.EncryptService
}



func NewEncrypt(
	app *fiber.App,
	encryptService domain.EncryptService,
){
	encryptAPI := EncryptAPI{
		encryptService:encryptService,
	}

	generate := app.Group("encrypt")

	generate.Post("/", encryptAPI.Create)
}

func (e *EncryptAPI) Create(ctx *fiber.Ctx) error {
	var req dto.CreateEncryptRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest, err.Error()))
	}

	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := e.encryptService.Encrypt(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Encrypt String", res))
}