package dto

//========== Customer Type ==========
type CustomerTypeRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description" binding:"required"`
}

type CreateOrUpdateCustomerTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CustomerTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
