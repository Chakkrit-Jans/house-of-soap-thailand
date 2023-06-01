package dto

//========== Customer Group ==========
type CustomerGroupRequest struct {
	Name        string `form:"name"  binding:"required"`
	Description string `form:"description" binding:"required"`
	VatID       uint   `form:"vatId"`
}

type CreateOrUpdateCustomerGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VatID       uint   `json:"vatId"`
}

type CustomerGroupResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Vat         VatResponse `json:"vat"`
}

type CustomerGroupOnlyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
