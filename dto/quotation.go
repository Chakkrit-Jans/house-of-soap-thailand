package dto

import (
	"hst-api/enum"
	"time"
)

// ========== Quotation ==========
type QuotationRequest struct {
	CustomersID        uint      `form:"customerId"  binding:"required"`
	DocumentDate       time.Time `form:"docdate" binding:"required"`
	DiscountPercentage float32   `form:"discountpercentage"`
	DiscountAmount     float32   `form:"discountamount"`
	SalesID            uint      `form:"salesId"  binding:"required"`
}

type CreateOrUpdateQuotationResponse struct {
	ID           uint      `json:"id"`
	Number       string    `json:"number"`
	Revision     uint8     `json:"revision"`
	CustomersID  uint      `json:"customerId"`
	CustomerCode string    `json:"customercode"`
	CustomerName string    `json:"customername"`
	VatID        uint      `json:"vatId"`
	DocumentDate time.Time `json:"docdate"`
	SalesID      uint      `json:"saleId"`
	SalesName    string    `json:"salesname"`

	Total              float32 `json:"total"`
	DiscountPercentage float32 `json:"discountpercentage"`
	DiscountAmount     float32 `json:"discountamount"`
	VatPercentage      float32 `form:"vatpercentage"`
	VatAmount          float32 `json:"vatamount"`
	GrandTotal         float32 `json:"grandtotal"`

	Status enum.QuotationStatus `json:"status"`
}

type QuotationResponse struct {
	ID           uint      `json:"id"`
	Number       string    `json:"number"`
	Revision     uint8     `json:"revision"`
	CustomersID  uint      `json:"customerId"`
	CustomerCode string    `json:"customercode"`
	CustomerName string    `json:"customername"`
	VatID        uint      `json:"vatId"`
	DocumentDate time.Time `json:"docdate"`
	SalesID      uint      `json:"saleId"`
	SalesName    string    `json:"salesname"`

	Total              float32 `json:"total"`
	DiscountPercentage float32 `json:"discountpercentage"`
	DiscountAmount     float32 `json:"discountamount"`
	VatPercentage      float32 `form:"vatpercentage"`
	VatAmount          float32 `json:"vatamount"`
	GrandTotal         float32 `json:"grandtotal"`

	Status enum.QuotationStatus `json:"status"`

	QuotationLine []QuotationLineResponse `json:"quotationline"`
}
