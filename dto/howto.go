package dto

type HowtoRequest struct {
	Code            string `form:"code"  binding:"required"`
	Description     string `form:"description"`
	DescriptionThai string `form:"descriptionthai"`
}

type CreateOrUpdateHowtoResponse struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	Description     string `json:"description"`
	DescriptionThai string `json:"descriptionthai"`
}

type HowtoResponse struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	Description     string `json:"description"`
	DescriptionThai string `json:"descriptionthai"`
}
