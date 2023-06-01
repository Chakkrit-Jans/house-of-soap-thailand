package dto

//========== Customer ==========
type CustomerRequest struct {
	Code      string `form:"code"  binding:"required"`
	FirstName string `form:"firstname" binding:"required"`
	LastName  string `form:"lastname"`
	FullName  string `form:"fullname" binding:"required"`
	NickName  string `form:"nickname"`
	Email     string `form:"email" binding:"required"`
	Mobile    string `form:"mobile"`
	Image     string `form:"image"`
	Status    uint8  `form:"status" binding:"required"`

	TitleID         uint `form:"titleId"`
	CustomerGroupID uint `form:"customergroupId"`
	CustomerTypeID  uint `form:"customertypeId"`
	VatID           uint `form:"vatId"`
}

type CreateOrUpdateCustomerResponse struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	FullName  string `json:"fullname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Image     string `json:"image"`
	Status    string `json:"status"`

	TitleID         uint `json:"titleId"`
	CustomerGroupID uint `json:"customergroupId"`
	CustomerTypeID  uint `json:"customertypeId"`
	VatID           uint `json:"vatId"`
}

type CustomerResponse struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	FullName  string `json:"fullname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Image     string `json:"image"`
	Status    uint8  `json:"status"`

	Title         TitleResponse             `json:"title"`
	CustomerGroup CustomerGroupOnlyResponse `json:"customergroup"`
	CustomerType  CustomerTypeResponse      `json:"customertype"`
	Vat           VatResponse               `json:"vat"`
}
