package dto

import "hst-api/enum"

//========== Quotation Line Option ==========
type QuotationLineOptionRequest struct {
	QuotationLineID uint `form:"quotationlineId"  binding:"required"`
	CostID          uint `form:"costId" binding:"required"`
}

type CreateOrUpdateQuotationLineOptionResponse struct {
	ID           uint          `json:"id"`
	CostID       uint          `json:"costId"`
	OptionCode   string        `json:"optioncode"`
	OptionName   string        `json:"optionname"`
	CostType     enum.CostType `json:"costtype"`
	Qty          int           `json:"qty"`
	Price        float32       `json:"price"`
	Amount       float32       `json:"amount"`
	CostPerUnit  float32       `json:"costperunit"`
	CostAmount   float32       `json:"costamount"`
	Profit       float32       `json:"profit"`
	ProfitAmount float32       `json:"profitamount"`

	QuotationLineID uint `json:"quotationlineId"  binding:"required"`
}

type QuotationLineOptionResponse struct {
	ID           uint          `json:"id"`
	CostID       uint          `json:"costId"`
	OptionCode   string        `json:"optioncode"`
	OptionName   string        `json:"optionname"`
	CostType     enum.CostType `json:"costtype"`
	Qty          int           `json:"qty"`
	Price        float32       `json:"price"`
	Amount       float32       `json:"amount"`
	CostPerUnit  float32       `json:"costperunit"`
	CostAmount   float32       `json:"costamount"`
	Profit       float32       `json:"profit"`
	ProfitAmount float32       `json:"profitamount"`
}
