package dto

type EncryptResponse struct{
	XorEncrypt string `json:"xor_encrypt"`
}


type CreateEncryptRequest struct {
	Key         string  `json:"key" validate:"required"`
	PlainText  string  `json:"plain_text" validate:"required"`                              
}