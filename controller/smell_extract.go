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

type SmellExtracts struct{}

func (s SmellExtracts) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var smellextracts []model.SmellExtracts

	if search != "" {
		db.Conn.Find(&smellextracts, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&smellextracts)
	}

	var result []dto.SmellExtractsResponse
	for _, smellextract := range smellextracts {
		result = append(result, dto.SmellExtractsResponse{
			ID:          smellextract.ID,
			Name:        smellextract.Name,
			Description: smellextract.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (s SmellExtracts) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var smellextract model.SmellExtracts

	if err := db.Conn.First(&smellextract, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SmellExtractsResponse{
		ID:          smellextract.ID,
		Name:        smellextract.Name,
		Description: smellextract.Description,
	})
}

func (s SmellExtracts) Create(ctx *gin.Context) {

	var form dto.SmellExtractsRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	smellextract := model.SmellExtracts{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&smellextract).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSmellExtractsResponse{
		ID:          smellextract.ID,
		Name:        smellextract.Name,
		Description: smellextract.Description,
	})

}

func (s SmellExtracts) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SmellExtractsRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var smellextract model.SmellExtracts
	if err := db.Conn.First(&smellextract, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	smellextract.Name = form.Name
	smellextract.Description = form.Description

	if err := db.Conn.Save(&smellextract).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSmellExtractsResponse{
		ID:          smellextract.ID,
		Name:        smellextract.Name,
		Description: smellextract.Description,
	})
}

func (s SmellExtracts) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var smellextracts model.SmellExtracts

	if err := db.Conn.First(&smellextracts, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&smellextracts, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&smellextracts, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
