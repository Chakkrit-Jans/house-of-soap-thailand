package dto

// ========== Product Type ==========
type ProductTypeRequest struct {
	Name string `form:"name" binding:"required"`

	ProductID uint `form:"productId"`
}

type CreateOrUpdateProductTypeResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	ProductID uint `json:"productId"`
}

type ProductTypeResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	Product ProductResponse `json:"product"`
}

type ProductTypeOnlyResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
