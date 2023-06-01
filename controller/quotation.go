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

type Quotation struct{}

func (q Quotation) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var quotations []model.Quotation

	query := db.Conn.Preload("QuotationLine").Preload("QuotationLine.QuotationLineOption")

	if search != "" {
		query.Find(&quotations, "number LIKE ? OR customer_code LIKE ? OR customer_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		query.Find(&quotations)
	}

	var result []dto.QuotationResponse

	for _, quotation := range quotations {
		quotationResult := dto.QuotationResponse{
			ID:           quotation.ID,
			Number:       quotation.Number,
			Revision:     quotation.Revision,
			CustomersID:  quotation.CustomersID,
			CustomerCode: quotation.CustomerCode,
			CustomerName: quotation.CustomerName,
			VatID:        quotation.VatID,
			DocumentDate: quotation.DocumentDate,
			SalesID:      quotation.SalesID,
			SalesName:    quotation.SalesName,
			Status:       quotation.Status,

			Total:              quotation.Total,
			DiscountPercentage: quotation.DiscountPercentage,
			DiscountAmount:     quotation.DiscountAmount,
			VatAmount:          quotation.VatAmount,
			GrandTotal:         quotation.GrandTotal,
		}

		var quotationlines []dto.QuotationLineResponse

		for _, quotationline := range quotation.QuotationLine {

			quotationlineResult := dto.QuotationLineResponse{
				ID:                    quotationline.ID,
				LineNum:               quotationline.LineNum,
				ProductFormulaID:      quotationline.ProductFormulaID,
				ItemCode:              quotationline.ItemCode,
				ItemName:              quotationline.ItemName,
				ProductTypeID:         quotationline.ProductTypeID,
				ProductTypeName:       quotationline.ProductTypeName,
				TestSizeID:            quotationline.TestSizeID,
				TestSizeName:          quotationline.TestSizeName,
				Qty:                   quotationline.Qty,
				QtyTestSize:           quotationline.QtyTestSize,
				Price:                 quotationline.Price,
				Amount:                quotationline.Amount,
				ItemCost:              quotationline.ItemCost,
				ProductionCostPerItem: quotationline.ProductionCostPerItem,
				ProductionCostPerUnit: quotationline.ProductionCostPerUnit,
				ProductionCostTotal:   quotationline.ProductionCostTotal,
				AdditionAmount:        quotationline.AdditionAmount,
				Profit:                quotationline.Profit,
				ProfitAmount:          quotationline.ProfitAmount,
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
		result = append(result, quotationResult)
	}

	ctx.JSON(http.StatusOK, result)

}

func (q Quotation) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var quotation model.Quotation

	query := db.Conn.Preload("QuotationLine").Preload("QuotationLine.QuotationLineOption")

	if err := query.First(&quotation, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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
		SalesID:      quotation.SalesID,
		SalesName:    quotation.SalesName,
		Status:       quotation.Status,

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

func (q Quotation) Create(ctx *gin.Context) {

	var form dto.QuotationRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var status enum.QuotationStatus = "Draft"

	quotation := model.Quotation{
		CustomersID:        form.CustomersID,
		DocumentDate:       form.DocumentDate,
		DiscountPercentage: form.DiscountPercentage,
		DiscountAmount:     form.DiscountAmount,
		SalesID:            form.SalesID,
		Status:             status,
	}

	if err := db.Conn.Create(&quotation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := quotation.ID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine").Preload("QuotationLine.QuotationLineOption")

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
		SalesID:      quotation_output.SalesID,
		SalesName:    quotation_output.SalesName,
		Status:       quotation.Status,

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

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)

}

func (q Quotation) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.QuotationRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var quotation model.Quotation
	if err := db.Conn.First(&quotation, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	quotation.CustomersID = form.CustomersID
	quotation.DocumentDate = form.DocumentDate
	quotation.DiscountPercentage = form.DiscountPercentage
	quotation.DiscountAmount = form.DiscountAmount
	quotation.SalesID = form.SalesID

	if err := db.Conn.Save(&quotation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//id := quotation.ID
	var quotation_output model.Quotation

	query := db.Conn.Preload("QuotationLine").Preload("QuotationLine.QuotationLineOption")

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
		SalesID:      quotation_output.SalesID,
		SalesName:    quotation_output.SalesName,
		Status:       quotation.Status,

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

	quotationResult.QuotationLine = quotationlines

	ctx.JSON(http.StatusOK, quotationResult)

}

func (q Quotation) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var quotation model.Quotation

	if err := db.Conn.First(&quotation, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&quotation, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
