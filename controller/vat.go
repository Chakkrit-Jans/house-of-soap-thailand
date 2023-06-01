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

type Vat struct{}

func (c Vat) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var vats []model.Vat

	if search != "" {
		db.Conn.Find(&vats, "name LIKE ?", "%"+search+"%")
	} else {
		db.Conn.Find(&vats)
	}

	var result []dto.VatResponse
	for _, vat := range vats {
		result = append(result, dto.VatResponse{
			ID:         vat.ID,
			Name:       vat.Name,
			Percentage: vat.Percentage,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Vat) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var vat model.Vat

	if err := db.Conn.First(&vat, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.VatResponse{
		ID:         vat.ID,
		Name:       vat.Name,
		Percentage: vat.Percentage,
	})
}

func (c Vat) Create(ctx *gin.Context) {

	var form dto.VatRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vat := model.Vat{
		Name:       form.Name,
		Percentage: form.Percentage,
	}

	if err := db.Conn.Create(&vat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateVatResponse{
		ID:         vat.ID,
		Name:       vat.Name,
		Percentage: vat.Percentage,
	})

}

func (c Vat) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.VatRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vat model.Vat
	if err := db.Conn.First(&vat, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	vat.Name = form.Name
	vat.Percentage = form.Percentage

	if err := db.Conn.Save(&vat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateVatResponse{
		ID:         vat.ID,
		Name:       vat.Name,
		Percentage: vat.Percentage,
	})
}

func (c Vat) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var vat model.Vat

	if err := db.Conn.First(&vat, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&vat, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&vat, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
