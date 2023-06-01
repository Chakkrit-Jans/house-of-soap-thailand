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

type Color3 struct{}

func (c Color3) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var colors3 []model.Color3

	if search != "" {
		db.Conn.Find(&colors3, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&colors3)
	}

	var result []dto.Color3Response
	for _, color3 := range colors3 {
		result = append(result, dto.Color3Response{
			ID:          color3.ID,
			Name:        color3.Name,
			Description: color3.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Color3) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var colors3 model.Color3

	if err := db.Conn.First(&colors3, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.Color3Response{
		ID:          colors3.ID,
		Name:        colors3.Name,
		Description: colors3.Description,
	})
}

func (c Color3) Create(ctx *gin.Context) {

	var form dto.Color3Request
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	color3 := model.Color3{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&color3).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateColor3Response{
		ID:          color3.ID,
		Name:        form.Name,
		Description: form.Description,
	})

}

func (c Color3) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.Color3Request

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var colors3 model.Color3
	if err := db.Conn.First(&colors3, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	colors3.Name = form.Name
	colors3.Description = form.Description

	if err := db.Conn.Save(&colors3).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateColor3Response{
		ID:          colors3.ID,
		Name:        colors3.Name,
		Description: colors3.Description,
	})
}

func (c Color3) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var color model.Color3

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
