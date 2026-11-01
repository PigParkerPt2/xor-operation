package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"xor-operation/domain"
	"xor-operation/dto"
)

type EncryptService struct {

}
func NewEncrypt() domain.EncryptService{
	return &EncryptService{}
}

func (es *EncryptService) Encrypt(ctx context.Context, req dto.CreateEncryptRequest) (dto.EncryptResponse, error) {
	inputBytes := []byte(req.PlainText)
	keyBytes := []byte(req.Key)

	if len(keyBytes) == 0 {
		return dto.EncryptResponse{}, errors.New("Key tidak boleh kosong")
	}

	var encrypted []byte
	for i, b := range inputBytes {
		keyByte := keyBytes[i%len(keyBytes)]

		encrypted = append(encrypted, b^keyByte)
	}

	var result []string
	for _, b := range encrypted {
		result = append(result, strconv.Itoa(int(b)))
	}
	return dto.EncryptResponse{XorEncrypt: strings.Join(result, " ")}, nil
}