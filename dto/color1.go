package dto

//========== Color1 ==========
type Color1Request struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`
}

type CreateOrUpdateColor1Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Color1Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
