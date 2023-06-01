package dto

type TitleRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`
}

type CreateOrUpdateTitleResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TitleResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
