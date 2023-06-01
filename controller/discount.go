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

type Discount struct{}

func (c Discount) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var discounts []model.Discount

	if search != "" {
		db.Conn.Find(&discounts, "Name LIKE ?", "%"+search+"%")
	} else {
		db.Conn.Find(&discounts)
	}

	var result []dto.DiscountResponse
	for _, discount := range discounts {
		result = append(result, dto.DiscountResponse{
			ID:                  discount.ID,
			Name:                discount.Name,
			PercentageDiscount:  discount.PercentageDiscount,
			LimitAmountDiscount: discount.LimitAmountDiscount,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Discount) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var discount model.Discount

	if err := db.Conn.First(&discount, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.DiscountResponse{
		ID:                  discount.ID,
		Name:                discount.Name,
		PercentageDiscount:  discount.PercentageDiscount,
		LimitAmountDiscount: discount.LimitAmountDiscount,
	})
}

func (c Discount) Create(ctx *gin.Context) {

	var form dto.DiscountRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discount := model.Discount{
		Name:                form.Name,
		PercentageDiscount:  form.PercentageDiscount,
		LimitAmountDiscount: form.LimitAmountDiscount,
	}

	if err := db.Conn.Create(&discount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateDiscountResponse{
		ID:                  discount.ID,
		Name:                discount.Name,
		PercentageDiscount:  discount.PercentageDiscount,
		LimitAmountDiscount: discount.LimitAmountDiscount,
	})

}

func (c Discount) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.DiscountRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var discount model.Discount
	if err := db.Conn.First(&discount, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	discount.Name = form.Name
	discount.PercentageDiscount = form.PercentageDiscount
	discount.LimitAmountDiscount = form.LimitAmountDiscount

	if err := db.Conn.Save(&discount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateDiscountResponse{
		ID:                  discount.ID,
		Name:                discount.Name,
		PercentageDiscount:  discount.PercentageDiscount,
		LimitAmountDiscount: discount.LimitAmountDiscount,
	})
}

func (c Discount) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var discount model.Discount

	if err := db.Conn.First(&discount, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&discount, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&discount, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
