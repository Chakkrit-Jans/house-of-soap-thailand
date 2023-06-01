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

type CustomerGroup struct{}

func (c CustomerGroup) FindAll(ctx *gin.Context) {

	vatId := ctx.Query("vatId")
	search := ctx.Query("search")

	var customergroups []model.CustomerGroup
	query := db.Conn.Preload("Vat")

	if vatId != "" {
		query = query.Where("vat_id = ?", vatId)
	}

	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Find(&customergroups)

	var result []dto.CustomerGroupResponse
	for _, customergroup := range customergroups {
		result = append(result, dto.CustomerGroupResponse{
			ID:          customergroup.ID,
			Name:        customergroup.Name,
			Description: customergroup.Description,
			Vat: dto.VatResponse{
				ID:         customergroup.Vat.ID,
				Name:       customergroup.Vat.Name,
				Percentage: customergroup.Vat.Percentage,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)

}

func (c CustomerGroup) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var customergroup model.CustomerGroup

	if err := db.Conn.First(&customergroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CustomerGroupResponse{
		ID:          customergroup.ID,
		Name:        customergroup.Name,
		Description: customergroup.Description,
		Vat: dto.VatResponse{
			ID:         customergroup.Vat.ID,
			Name:       customergroup.Vat.Name,
			Percentage: customergroup.Vat.Percentage,
		},
	})
}

func (c CustomerGroup) Create(ctx *gin.Context) {

	var form dto.CustomerGroupRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customergroup := model.CustomerGroup{
		Name:        form.Name,
		Description: form.Description,
		VatID:       form.VatID,
	}

	if err := db.Conn.Create(&customergroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateCustomerGroupResponse{
		ID:          customergroup.ID,
		Name:        customergroup.Name,
		Description: customergroup.Description,
		VatID:       customergroup.VatID,
	})

}

func (c CustomerGroup) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.CustomerGroupRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customergroup model.CustomerGroup
	if err := db.Conn.First(&customergroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	customergroup.Name = form.Name
	customergroup.Description = form.Description
	customergroup.VatID = form.VatID

	if err := db.Conn.Save(&customergroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateCustomerGroupResponse{
		ID:          customergroup.ID,
		Name:        customergroup.Name,
		Description: customergroup.Description,
		VatID:       customergroup.VatID,
	})
}

func (c CustomerGroup) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var customergroup model.CustomerGroup

	if err := db.Conn.First(&customergroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&customergroup, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&customergroup, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
