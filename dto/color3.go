package dto

//========== Color3 ==========
type Color3Request struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`
}

type CreateOrUpdateColor3Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Color3Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
