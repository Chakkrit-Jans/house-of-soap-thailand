package dto

type VatRequest struct {
	Name       string  `form:"name" binding:"required"`
	Percentage float32 `form:"percentage"`
}

type CreateOrUpdateVatResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Percentage float32 `json:"percentage"`
}

type VatResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Percentage float32 `json:"percentage"`
}
