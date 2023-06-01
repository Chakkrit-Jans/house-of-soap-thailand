package controller

import (
	"errors"
	"hst-api/db"
	"hst-api/dto"
	"hst-api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuotationLine struct{}

func (q QuotationLine) FindLineAll(ctx *gin.Context) {

	quotation_id := ctx.Param("quotationId")
	// lineid := ctx.Param("lineId")
	var quotation model.Quotation

	query := db.Conn.Preload("QuotationLine").Preload("QuotationLine.QuotationLineOption")

	if err := query.Where("number = ?", quotation_id).First(&quotation).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotationResult := dto.QuotationResponse{
		ID:           quotation.ID,
		Number:       quotation.Number,
		Revision:     quotation.Revision,
		CustomersID:  quotation.CustomersID,
		CustomerCode: quotation.CustomerCode,
		CustomerName: quotation.CustomerName,
		VatID:        quotation.VatID,
		DocumentDate: quotation.DocumentDate,

		Total:              quotation.Total,
		DiscountPercentage: quotation.DiscountPercentage,
		DiscountAmount:     quotation.DiscountAmount,
		VatAmount:          quotation.VatAmount,
		GrandTotal:         quotation.GrandTotal,
	}

	var quotationlines []dto.QuotationLineResponse

	for _, quotationline := range quotation.QuotationLine {

		quotationlineResult := dto.QuotationLineResponse{
			ID:               quotationline.ID,
			LineNum:          quotationline.LineNum,
			ProductFormulaID: quotationline.ProductFormulaID,
			ItemCode:         quotationline.ItemCode,
			ItemName:         quotationline.ItemName,
			ProductID:        quotationline.ProductID,
			ProductName:      quotationline.ProductName,
			ProductTypeID:    quotationline.ProductTypeID,
			ProductTypeName:  quotationline.ProductTypeName,
			TestSizeID:       quotationline.TestSizeID,
			TestSizeName:     quotationline.TestSizeName,
			Qty:              quotationline.Qty,
			QtyTestSize:      quotationline.QtyTestSize,
			Price:            quotationline.Price,
			Amount:           quotationline.Amount,

			//---------- Cost for calculation ----------
			ItemCost:              quotationline.ItemCost,
			ProductionCostPerItem: quotationline.ProductionCostPerItem,
			ProductionCostPerUnit: quotationline.ProductionCostPerUnit,
			ProductionCostTotal:   quotationline.ProductionCostTotal,
			AdditionAmount:        quotationline.AdditionAmount,

			//---------- Profit for calculation ----------
			Profit:       quotationline.Profit,
			ProfitAmount: quotationline.ProfitAmount,
		}

		var quotationlineoptions []dto.QuotationLineOptionResponse

		for _, quotationlineoption := range quotationline.QuotationLineOption {
			quotationlineoptions = append(quotationlineoptions, dto.QuotationLineOptionResponse{
				ID:         quotationlineoption.ID,
				CostID:     quotationlineoption.CostID,
				OptionCode: quotationlineoption.OptionCode,
				OptionName: quotationlineoption.OptionName,
				CostType:   quotationlineoption.CostType,
				Qty:        quotationlineoption.Qty,
				Price:      quotationlineoption.Price,
				Amount:     quotationlineoption.Amount,
			})
		}
		quotationlineResult.QuotationLineOption = quotationlineoptions
		quotationlines = append(quotationlines, quotationlineResult)
	}

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)
}

func (q QuotationLine) FindLineOne(ctx *gin.Context) {

	quotation_id := ctx.Param("quotationId")
	line_id := ctx.Param("lineId")
	var quotation model.Quotation

	query := db.Conn.Preload("QuotationLine", line_id).Preload("QuotationLine.QuotationLineOption")
	//query := db.Conn.Joins("QuotationLine", line_id).Joins("QuotationLine.QuotationLineOption")

	if err := query.Where("number = ?", quotation_id).First(&quotation).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotationResult := dto.QuotationResponse{
		ID:           quotation.ID,
		Number:       quotation.Number,
		Revision:     quotation.Revision,
		CustomersID:  quotation.CustomersID,
		CustomerCode: quotation.CustomerCode,
		CustomerName: quotation.CustomerName,
		VatID:        quotation.VatID,
		DocumentDate: quotation.DocumentDate,

		Total:              quotation.Total,
		DiscountPercentage: quotation.DiscountPercentage,
		DiscountAmount:     quotation.DiscountAmount,
		VatAmount:          quotation.VatAmount,
		GrandTotal:         quotation.GrandTotal,
	}

	var quotationlines []dto.QuotationLineResponse

	for _, quotationline := range quotation.QuotationLine {

		quotationlineResult := dto.QuotationLineResponse{
			ID:               quotationline.ID,
			LineNum:          quotationline.LineNum,
			ProductFormulaID: quotationline.ProductFormulaID,
			ItemCode:         quotationline.ItemCode,
			ItemName:         quotationline.ItemName,
			ProductID:        quotationline.ProductID,
			ProductName:      quotationline.ProductName,
			ProductTypeID:    quotationline.ProductTypeID,
			ProductTypeName:  quotationline.ProductTypeName,
			TestSizeID:       quotationline.TestSizeID,
			TestSizeName:     quotationline.TestSizeName,
			Qty:              quotationline.Qty,
			QtyTestSize:      quotationline.QtyTestSize,
			Price:            quotationline.Price,
			Amount:           quotationline.Amount,

			//---------- Cost for calculation ----------
			ItemCost:              quotationline.ItemCost,
			ProductionCostPerItem: quotationline.ProductionCostPerItem,
			ProductionCostPerUnit: quotationline.ProductionCostPerUnit,
			ProductionCostTotal:   quotationline.ProductionCostTotal,
			AdditionAmount:        quotationline.AdditionAmount,

			//---------- Profit for calculation ----------
			Profit:       quotationline.Profit,
			ProfitAmount: quotationline.ProfitAmount,
		}

		var quotationlineoptions []dto.QuotationLineOptionResponse

		for _, quotationlineoption := range quotationline.QuotationLineOption {
			quotationlineoptions = append(quotationlineoptions, dto.QuotationLineOptionResponse{
				ID:         quotationlineoption.ID,
				CostID:     quotationlineoption.CostID,
				OptionCode: quotationlineoption.OptionCode,
				OptionName: quotationlineoption.OptionName,
				CostType:   quotationlineoption.CostType,
				Qty:        quotationlineoption.Qty,
				Price:      quotationlineoption.Price,
				Amount:     quotationlineoption.Amount,
			})
		}
		quotationlineResult.QuotationLineOption = quotationlineoptions
		quotationlines = append(quotationlines, quotationlineResult)
	}

	if quotationlines == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)
}

func (q QuotationLine) Create(ctx *gin.Context) {

	var form dto.QuotationLineRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quotationline := model.QuotationLine{
		ProductFormulaID: form.ProductFormulaID,
		ProductID:        form.ProductID,
		ProductTypeID:    form.ProductTypeID,
		TestSizeID:       form.TestSizeID,
		Qty:              form.Qty,
		QuotationID:      form.QuotationID,
	}

	if err := db.Conn.Create(&quotationline).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := quotationline.QuotationID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine", quotationline.ID).Preload("QuotationLine.QuotationLineOption")

	if err := query.First(&quotation_output, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotationResult := dto.QuotationResponse{
		ID:           quotation_output.ID,
		Number:       quotation_output.Number,
		Revision:     quotation_output.Revision,
		CustomersID:  quotation_output.CustomersID,
		CustomerCode: quotation_output.CustomerCode,
		CustomerName: quotation_output.CustomerName,
		VatID:        quotation_output.VatID,
		DocumentDate: quotation_output.DocumentDate,

		Total:              quotation_output.Total,
		DiscountPercentage: quotation_output.DiscountPercentage,
		DiscountAmount:     quotation_output.DiscountAmount,
		VatAmount:          quotation_output.VatAmount,
		GrandTotal:         quotation_output.GrandTotal,
	}

	var quotationlines []dto.QuotationLineResponse

	for _, quotationline := range quotation_output.QuotationLine {

		quotationlineResult := dto.QuotationLineResponse{
			ID:               quotationline.ID,
			LineNum:          quotationline.LineNum,
			ProductFormulaID: quotationline.ProductFormulaID,
			ItemCode:         quotationline.ItemCode,
			ItemName:         quotationline.ItemName,
			ProductID:        quotationline.ProductID,
			ProductName:      quotationline.ProductName,
			ProductTypeID:    quotationline.ProductTypeID,
			ProductTypeName:  quotationline.ProductTypeName,
			TestSizeID:       quotationline.TestSizeID,
			TestSizeName:     quotationline.TestSizeName,
			Qty:              quotationline.Qty,
			QtyTestSize:      quotationline.QtyTestSize,
			Price:            quotationline.Price,
			Amount:           quotationline.Amount,

			//---------- Cost for calculation ----------
			ItemCost:              quotationline.ItemCost,
			ProductionCostPerItem: quotationline.ProductionCostPerItem,
			ProductionCostPerUnit: quotationline.ProductionCostPerUnit,
			ProductionCostTotal:   quotationline.ProductionCostTotal,
			AdditionAmount:        quotationline.AdditionAmount,

			//---------- Profit for calculation ----------
			Profit:       quotationline.Profit,
			ProfitAmount: quotationline.ProfitAmount,
		}

		var quotationlineoptions []dto.QuotationLineOptionResponse

		for _, quotationlineoption := range quotationline.QuotationLineOption {
			quotationlineoptions = append(quotationlineoptions, dto.QuotationLineOptionResponse{
				ID:         quotationlineoption.ID,
				CostID:     quotationlineoption.CostID,
				OptionCode: quotationlineoption.OptionCode,
				OptionName: quotationlineoption.OptionName,
				CostType:   quotationlineoption.CostType,
				Qty:        quotationlineoption.Qty,
				Price:      quotationlineoption.Price,
				Amount:     quotationlineoption.Amount,
			})
		}
		quotationlineResult.QuotationLineOption = quotationlineoptions
		quotationlines = append(quotationlines, quotationlineResult)
	}

	if quotationlines == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)

}

func (q QuotationLine) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.QuotationLineRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var quotationline model.QuotationLine
	if err := db.Conn.First(&quotationline, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotationline.ProductFormulaID = form.ProductFormulaID
	quotationline.ProductID = form.ProductID
	quotationline.ProductTypeID = form.ProductTypeID
	quotationline.TestSizeID = form.TestSizeID
	quotationline.Qty = form.Qty
	quotationline.QuotationID = form.QuotationID

	if err := db.Conn.Save(&quotationline).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//id := quotationline.QuotationID
	quotation_id := quotationline.QuotationID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine", quotationline.ID).Preload("QuotationLine.QuotationLineOption")

	if err := query.First(&quotation_output, quotation_id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotationResult := dto.QuotationResponse{
		ID:           quotation_output.ID,
		Number:       quotation_output.Number,
		Revision:     quotation_output.Revision,
		CustomersID:  quotation_output.CustomersID,
		CustomerCode: quotation_output.CustomerCode,
		CustomerName: quotation_output.CustomerName,
		VatID:        quotation_output.VatID,
		DocumentDate: quotation_output.DocumentDate,

		Total:              quotation_output.Total,
		DiscountPercentage: quotation_output.DiscountPercentage,
		DiscountAmount:     quotation_output.DiscountAmount,
		VatAmount:          quotation_output.VatAmount,
		GrandTotal:         quotation_output.GrandTotal,
	}

	var quotationlines []dto.QuotationLineResponse

	for _, quotationline := range quotation_output.QuotationLine {

		quotationlineResult := dto.QuotationLineResponse{
			ID:               quotationline.ID,
			LineNum:          quotationline.LineNum,
			ProductFormulaID: quotationline.ProductFormulaID,
			ItemCode:         quotationline.ItemCode,
			ItemName:         quotationline.ItemName,
			ProductID:        quotationline.ProductID,
			ProductName:      quotationline.ProductName,
			ProductTypeID:    quotationline.ProductTypeID,
			ProductTypeName:  quotationline.ProductTypeName,
			TestSizeID:       quotationline.TestSizeID,
			TestSizeName:     quotationline.TestSizeName,
			Qty:              quotationline.Qty,
			QtyTestSize:      quotationline.QtyTestSize,
			Price:            quotationline.Price,
			Amount:           quotationline.Amount,

			//---------- Cost for calculation ----------
			ItemCost:              quotationline.ItemCost,
			ProductionCostPerItem: quotationline.ProductionCostPerItem,
			ProductionCostPerUnit: quotationline.ProductionCostPerUnit,
			ProductionCostTotal:   quotationline.ProductionCostTotal,
			AdditionAmount:        quotationline.AdditionAmount,

			//---------- Profit for calculation ----------
			Profit:       quotationline.Profit,
			ProfitAmount: quotationline.ProfitAmount,
		}

		var quotationlineoptions []dto.QuotationLineOptionResponse

		for _, quotationlineoption := range quotationline.QuotationLineOption {
			quotationlineoptions = append(quotationlineoptions, dto.QuotationLineOptionResponse{
				ID:         quotationlineoption.ID,
				CostID:     quotationlineoption.CostID,
				OptionCode: quotationlineoption.OptionCode,
				OptionName: quotationlineoption.OptionName,
				CostType:   quotationlineoption.CostType,
				Qty:        quotationlineoption.Qty,
				Price:      quotationlineoption.Price,
				Amount:     quotationlineoption.Amount,
			})
		}
		quotationlineResult.QuotationLineOption = quotationlineoptions
		quotationlines = append(quotationlines, quotationlineResult)
	}

	if quotationlines == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)

}

func (q QuotationLine) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var quotationline model.QuotationLine

	if err := db.Conn.First(&quotationline, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&quotationline, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
