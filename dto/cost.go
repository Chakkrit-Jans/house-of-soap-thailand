package dto

import "hst-api/enum"

// ========== Cost ==========
type CostRequest struct {
	Code      string         `form:"code"  binding:"required"`
	Name      string         `form:"name"  binding:"required"`
	CostType  enum.CostType  `form:"costtype" binding:"required"`
	CostGroup enum.CostGroup `form:"costgroup" binding:"required"`
	Amount    float32        `form:"amount" binding:"required"`
}

type CreateOrUpdateCostResponse struct {
	ID        uint           `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CostType  enum.CostType  `json:"costtype"`
	CostGroup enum.CostGroup `json:"costgroup"`
	Amount    float32        `json:"amount"`
}

type CostResponse struct {
	ID        uint           `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CostType  enum.CostType  `json:"costtype"`
	CostGroup enum.CostGroup `json:"costgroup"`
	Amount    float32        `json:"amount"`
}
