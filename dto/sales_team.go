package dto

//========== Sales Team ==========
type SalesTeamRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description"`
}

type CreateOrUpdateSalesTeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SalesTeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
