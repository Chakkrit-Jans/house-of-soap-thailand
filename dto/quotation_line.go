package dto

//========== Quotation Line ==========
type QuotationLineRequest struct {
	// LineNum          int     `form:"linenum" binding:"required"`
	ProductFormulaID uint `form:"itemId" binding:"required"`
	ProductID        uint `form:"productId" binding:"required"`
	ProductTypeID    uint `form:"producttypeId" binding:"required"`
	TestSizeID       uint `form:"testsizeId" binding:"required"`
	Qty              int  `form:"qty"  binding:"required"`
	// Price            float32 `form:"price"`
	// Amount           float32 `form:"amount"`

	//---------- Quotation ----------
	QuotationID uint `form:"quotationId"`
}

type CreateOrUpdateQuotationLineResponse struct {
	ID               uint    `json:"id"`
	LineNum          int     `json:"line"`
	ProductFormulaID uint    `json:"itemId"`
	ItemCode         string  `json:"itemcode"`
	ItemName         string  `json:"itemname"`
	ProductID        uint    `json:"productId"`
	ProductName      string  `json:"productname"`
	ProductTypeID    uint    `json:"producttypeId"`
	ProductTypeName  string  `json:"producttypename"`
	TestSizeID       uint    `json:"testsizeId"`
	TestSizeName     string  `json:"testsizename"`
	Qty              int     `json:"qty"`
	QtyTestSize      int     `json:"qtyperunit"`
	Price            float32 `json:"price"`
	Amount           float32 `json:"amount"`

	//---------- Cost for calculation ----------
	ItemCost              float32 `json:"itemcost"`
	ProductionCostPerItem float32 `json:"productioncostperitem"`
	ProductionCostPerUnit float32 `json:"productioncostperunit"`
	ProductionCostTotal   float32 `json:"productioncosttotal"`
	AdditionAmount        float32 `json:"aditionadmount"`

	//---------- Profit for calculation ----------
	Profit       float32 `json:"profit"`
	ProfitAmount float32 `json:"profitamount"`

	//---------- Quotation ----------
	QuotationID uint `json:"quotationId"`
}

type QuotationLineResponse struct {
	ID               uint    `json:"id"`
	LineNum          int     `json:"line"`
	ProductFormulaID uint    `json:"itemId"`
	ItemCode         string  `json:"itemcode"`
	ItemName         string  `json:"itemname"`
	ProductID        uint    `json:"productId"`
	ProductName      string  `json:"productname"`
	ProductTypeID    uint    `json:"producttypeId"`
	ProductTypeName  string  `json:"producttypename"`
	TestSizeID       uint    `json:"testsizeId"`
	TestSizeName     string  `json:"testsizename"`
	Qty              int     `json:"qty"`
	QtyTestSize      int     `json:"qtyperunit"`
	Price            float32 `json:"price"`
	Amount           float32 `json:"amount"`

	//---------- Cost for calculation ----------
	ItemCost              float32 `json:"itemcost"`
	ProductionCostPerItem float32 `json:"productioncostperitem"`
	ProductionCostPerUnit float32 `json:"productioncostperunit"`
	ProductionCostTotal   float32 `json:"productioncosttotal"`
	AdditionAmount        float32 `json:"aditionadmount"`

	//---------- Profit for calculation ----------
	Profit       float32 `json:"profit"`
	ProfitAmount float32 `json:"profitamount"`

	QuotationLineOption []QuotationLineOptionResponse `json:"quotationlineoption"`
}
