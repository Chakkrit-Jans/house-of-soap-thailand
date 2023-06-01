package dto

type DiscountRequest struct {
	Name                string  `form:"name"  binding:"required"`
	PercentageDiscount  float32 `form:"percentagediscount" binding:"required"`
	LimitAmountDiscount float32 `form:"limitamountdiscount" binding:"required"`
}

type CreateOrUpdateDiscountResponse struct {
	ID                  uint    `json:"id"`
	Name                string  `json:"name"`
	PercentageDiscount  float32 `json:"percentagediscount"`
	LimitAmountDiscount float32 `json:"limitamountdiscount"`
}

type DiscountResponse struct {
	ID                  uint    `json:"id"`
	Name                string  `json:"name"`
	PercentageDiscount  float32 `json:"percentagediscount"`
	LimitAmountDiscount float32 `json:"limitamountdiscount"`
}
