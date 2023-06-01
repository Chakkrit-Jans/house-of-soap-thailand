package dto

//========== Color2 ==========
type Color2Request struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`
}

type CreateOrUpdateColor2Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Color2Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
