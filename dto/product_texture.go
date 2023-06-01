package dto

// ========== Product Texture ==========
type ProductTextureRequest struct {
	Code        string `form:"code" binding:"required"`
	Name        string `form:"name"`
	Description string `form:"description"`
}

type CreateOrUpdateProductTextureResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductTextureResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
