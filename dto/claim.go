package dto

type ClaimRequest struct {
	Code            string `form:"code"  binding:"required"`
	Description     string `form:"description"`
	DescriptionThai string `form:"descriptionthai"`
}

type CreateOrUpdateClaimResponse struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	Description     string `json:"description"`
	DescriptionThai string `json:"descriptionthai"`
}

type ClaimResponse struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	Description     string `json:"description"`
	DescriptionThai string `json:"descriptionthai"`
}
