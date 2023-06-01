package dto

// ========== Product ==========
type ProductRequest struct {
	Code string `form:"code"  binding:"required"`
	Name string `form:"name" binding:"required"`
}

type CreateOrUpdateProductResponse struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ProductResponse struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
