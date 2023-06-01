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

type SmellType struct{}

func (s SmellType) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var smelltypes []model.SmellType

	if search != "" {
		db.Conn.Find(&smelltypes, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&smelltypes)
	}

	var result []dto.SmellTypeResponse
	for _, smelltype := range smelltypes {
		result = append(result, dto.SmellTypeResponse{
			ID:          smelltype.ID,
			Name:        smelltype.Name,
			Description: smelltype.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (s SmellType) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var smelltype model.SmellType

	if err := db.Conn.First(&smelltype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SmellTypeResponse{
		ID:          smelltype.ID,
		Name:        smelltype.Name,
		Description: smelltype.Description,
	})
}

func (s SmellType) Create(ctx *gin.Context) {

	var form dto.SmellTypeRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	smelltype := model.SmellType{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&smelltype).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSmellTypeResponse{
		ID:          smelltype.ID,
		Name:        smelltype.Name,
		Description: smelltype.Description,
	})

}

func (s SmellType) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SmellTypeRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var smelltype model.SmellType
	if err := db.Conn.First(&smelltype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	smelltype.Name = form.Name
	smelltype.Description = form.Description

	if err := db.Conn.Save(&smelltype).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSmellTypeResponse{
		ID:          smelltype.ID,
		Name:        smelltype.Name,
		Description: smelltype.Description,
	})
}

func (s SmellType) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var smelltype model.SmellType

	if err := db.Conn.First(&smelltype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&smelltype, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&smellextracts, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
