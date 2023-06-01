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

type Color1 struct{}

func (c Color1) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var colors1 []model.Color1

	if search != "" {
		db.Conn.Find(&colors1, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&colors1)
	}

	var result []dto.Color1Response
	for _, color1 := range colors1 {
		result = append(result, dto.Color1Response{
			ID:          color1.ID,
			Name:        color1.Name,
			Description: color1.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Color1) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var colors1 model.Color1

	if err := db.Conn.First(&colors1, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.Color1Response{
		ID:          colors1.ID,
		Name:        colors1.Name,
		Description: colors1.Description,
	})
}

func (c Color1) Create(ctx *gin.Context) {

	var form dto.Color1Request
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	color1 := model.Color1{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&color1).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateColor1Response{
		ID:          color1.ID,
		Name:        form.Name,
		Description: form.Description,
	})

}

func (c Color1) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.Color1Request

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var colors1 model.Color1
	if err := db.Conn.First(&colors1, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	colors1.Name = form.Name
	colors1.Description = form.Description

	if err := db.Conn.Save(&colors1).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateColor1Response{
		ID:          colors1.ID,
		Name:        colors1.Name,
		Description: colors1.Description,
	})
}

func (c Color1) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var color model.Color1

	if err := db.Conn.First(&color, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&color, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&color, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
