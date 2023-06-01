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

type CustomerType struct{}

func (c CustomerType) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var CustomerTypes []model.CustomerType

	if search != "" {
		db.Conn.Find(&CustomerTypes, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&CustomerTypes)
	}

	var result []dto.CustomerTypeResponse
	for _, CustomerType := range CustomerTypes {
		result = append(result, dto.CustomerTypeResponse{
			ID:          CustomerType.ID,
			Name:        CustomerType.Name,
			Description: CustomerType.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c CustomerType) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var CustomerType model.CustomerType

	if err := db.Conn.First(&CustomerType, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CustomerTypeResponse{
		ID:          CustomerType.ID,
		Name:        CustomerType.Name,
		Description: CustomerType.Description,
	})
}

func (c CustomerType) Create(ctx *gin.Context) {

	var form dto.CustomerTypeRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customertype := model.CustomerType{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&customertype).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateCustomerTypeResponse{
		ID:          customertype.ID,
		Name:        customertype.Name,
		Description: customertype.Description,
	})

}

func (c CustomerType) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.CustomerTypeRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customertype model.CustomerType
	if err := db.Conn.First(&customertype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	customertype.Name = form.Name
	customertype.Description = form.Description

	if err := db.Conn.Save(&customertype).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateCustomerTypeResponse{
		ID:          customertype.ID,
		Name:        customertype.Name,
		Description: customertype.Description,
	})
}

func (c CustomerType) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var customertype model.CustomerType

	if err := db.Conn.First(&customertype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&customertype, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&customertype, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
