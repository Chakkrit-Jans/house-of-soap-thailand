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

type Color2 struct{}

func (c Color2) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var colors2 []model.Color2

	if search != "" {
		db.Conn.Find(&colors2, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&colors2)
	}

	var result []dto.Color2Response
	for _, color2 := range colors2 {
		result = append(result, dto.Color2Response{
			ID:          color2.ID,
			Name:        color2.Name,
			Description: color2.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Color2) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var colors2 model.Color2

	if err := db.Conn.First(&colors2, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.Color2Response{
		ID:          colors2.ID,
		Name:        colors2.Name,
		Description: colors2.Description,
	})
}

func (c Color2) Create(ctx *gin.Context) {

	var form dto.Color2Request
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	color2 := model.Color2{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&color2).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateColor2Response{
		ID:          color2.ID,
		Name:        form.Name,
		Description: form.Description,
	})

}

func (c Color2) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.Color2Request

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var colors2 model.Color2
	if err := db.Conn.First(&colors2, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	colors2.Name = form.Name
	colors2.Description = form.Description

	if err := db.Conn.Save(&colors2).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateColor2Response{
		ID:          colors2.ID,
		Name:        colors2.Name,
		Description: colors2.Description,
	})
}

func (c Color2) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var color model.Color2

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
