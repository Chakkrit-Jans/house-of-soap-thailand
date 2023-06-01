package controller

import (
	"errors"
	"hst-api/db"
	"hst-api/dto"
	"hst-api/enum"
	"hst-api/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuotationLineOption struct{}

func (q QuotationLineOption) FindLineOptionAll(ctx *gin.Context) {

	quotation_id := ctx.Param("quotationId")
	line_id := ctx.Param("lineId")
	var quotation model.Quotation

	query := db.Conn.Preload("QuotationLine", line_id).Preload("QuotationLine.QuotationLineOption")

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

				CostPerUnit:  quotationlineoption.CostPerUnit,
				CostAmount:   quotationlineoption.CostAmount,
				Profit:       quotationlineoption.Profit,
				ProfitAmount: quotationlineoption.ProfitAmount,
			})
		}
		quotationlineResult.QuotationLineOption = quotationlineoptions
		quotationlines = append(quotationlines, quotationlineResult)
	}

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)
}

func (q QuotationLineOption) FindLineOptionOne(ctx *gin.Context) {

	quotation_id := ctx.Param("quotationId")
	line_id := ctx.Param("lineId")
	option_id := ctx.Param("optionId")
	var quotation model.Quotation

	query := db.Conn.Preload("QuotationLine", line_id).Preload("QuotationLine.QuotationLineOption", option_id)

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

				CostPerUnit:  quotationlineoption.CostPerUnit,
				CostAmount:   quotationlineoption.CostAmount,
				Profit:       quotationlineoption.Profit,
				ProfitAmount: quotationlineoption.ProfitAmount,
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

func (q QuotationLineOption) Create(ctx *gin.Context) {

	var form dto.QuotationLineOptionRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var costtype enum.CostType = ""

	quotationlineoption := model.QuotationLineOption{
		QuotationLineID: form.QuotationLineID,
		CostID:          form.CostID,
		CostType:        costtype,
	}

	if err := db.Conn.Create(&quotationlineoption).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*===== Find QuotationID =====*/
	var quotationline model.QuotationLine
	if err := db.Conn.First(&quotationline, quotationlineoption.QuotationLineID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	/*===== Find QuotationID =====*/

	id := quotationline.QuotationID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine", quotationlineoption.QuotationLineID).Preload("QuotationLine.QuotationLineOption")

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

				CostPerUnit:  quotationlineoption.CostPerUnit,
				CostAmount:   quotationlineoption.CostAmount,
				Profit:       quotationlineoption.Profit,
				ProfitAmount: quotationlineoption.ProfitAmount,
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

func (q QuotationLineOption) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.QuotationLineOptionRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var quotationlineoption model.QuotationLineOption
	if err := db.Conn.First(&quotationlineoption, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var costtype enum.CostType = ""

	quotationlineoption.QuotationLineID = form.QuotationLineID
	quotationlineoption.CostID = form.CostID
	quotationlineoption.CostType = costtype

	if err := db.Conn.Save(&quotationlineoption).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*===== Find QuotationID =====*/
	var quotationline model.QuotationLine
	if err := db.Conn.First(&quotationline, quotationlineoption.QuotationLineID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	/*===== Find QuotationID =====*/

	quotationid := quotationline.QuotationID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine", quotationlineoption.QuotationLineID).Preload("QuotationLine.QuotationLineOption")

	if err := query.First(&quotation_output, quotationid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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

				CostPerUnit:  quotationlineoption.CostPerUnit,
				CostAmount:   quotationlineoption.CostAmount,
				Profit:       quotationlineoption.Profit,
				ProfitAmount: quotationlineoption.ProfitAmount,
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

func (q QuotationLineOption) Delete(ctx *gin.Context) {

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
