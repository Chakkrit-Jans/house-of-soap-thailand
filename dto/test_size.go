package dto

type TestSizeRequest struct {
	Name    string `form:"name"  binding:"required"`
	Qty     int    `form:"qty" binding:"required"`
	UnitsID uint   `form:"unitId"`
}

type CreateOrUpdateTestSizeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`

	UnitsID uint `json:"unitId"`
}

type TestSizeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`

	Units UnitResponse `json:"units"`
}
