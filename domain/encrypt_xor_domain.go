package domain

import (
	"context"
	"xor-operation/dto"
)

type EncryptStruct struct {
	Key string `json:"key"`
}

type EncryptService interface {
	Encrypt(ctx context.Context, req dto.CreateEncryptRequest) (dto.EncryptResponse, error)
}