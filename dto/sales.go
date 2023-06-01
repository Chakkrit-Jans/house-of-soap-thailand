package dto

//========== Sales ==========
type SalesRequest struct {
	Code      string `form:"code"  binding:"required"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
	FullName  string `form:"fullname"`
	NickName  string `form:"nickname"`
	Email     string `form:"email"`
	Mobile    string `form:"mobile"`
	Status    uint8  `form:"status"`

	TitleID     uint `form:"titleId"`
	SalesTeamID uint `form:"salesteamId"`
}

type CreateOrUpdateSalesResponse struct {
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

	TitleID     uint `json:"titleId"`
	SalesTeamID uint `json:"salesteamId"`
}

type SalesResponse struct {
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

	Title     TitleResponse     `json:"title"`
	SalesTeam SalesTeamResponse `json:"salesteam"`
}

type SalesOnlyResponse struct {
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
}
