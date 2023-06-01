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

type Units struct{}

func (u Units) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var units []model.Units

	if search != "" {
		db.Conn.Find(&units, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&units)
	}

	var result []dto.UnitResponse
	for _, unit := range units {
		result = append(result, dto.UnitResponse{
			ID:          unit.ID,
			Name:        unit.Name,
			Description: unit.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (u Units) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var unit model.Units

	if err := db.Conn.First(&unit, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.UnitResponse{
		ID:          unit.ID,
		Name:        unit.Name,
		Description: unit.Description,
	})
}

func (u Units) Create(ctx *gin.Context) {

	var form dto.UnitRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit := model.Units{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&unit).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateUnitResponse{
		ID:          unit.ID,
		Name:        unit.Name,
		Description: unit.Description,
	})

}

func (u Units) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.UnitRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var unit model.Units
	if err := db.Conn.First(&unit, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	unit.Name = form.Name
	unit.Description = form.Description

	if err := db.Conn.Save(&unit).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateUnitResponse{
		ID:          unit.ID,
		Name:        unit.Name,
		Description: unit.Description,
	})
}

func (u Units) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var unit model.Units

	if err := db.Conn.First(&unit, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&unit, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&unit, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
