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

type ProductTexture struct{}

func (p ProductTexture) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var producttextures []model.ProductTexture

	if search != "" {
		db.Conn.Find(&producttextures, "code LIKE ? OR name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&producttextures)
	}

	var result []dto.ProductTextureResponse
	for _, producttexture := range producttextures {
		result = append(result, dto.ProductTextureResponse{
			ID:          producttexture.ID,
			Code:        producttexture.Code,
			Name:        producttexture.Name,
			Description: producttexture.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p ProductTexture) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var producttexture model.ProductTexture

	if err := db.Conn.First(&producttexture, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductTextureResponse{
		ID:          producttexture.ID,
		Code:        producttexture.Code,
		Name:        producttexture.Name,
		Description: producttexture.Description,
	})
}

func (p ProductTexture) Create(ctx *gin.Context) {

	var form dto.ProductTextureRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	producttexture := model.ProductTexture{
		Code:        form.Code,
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&producttexture).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateProductTextureResponse{
		ID:          producttexture.ID,
		Code:        producttexture.Code,
		Name:        producttexture.Name,
		Description: producttexture.Description,
	})

}

func (p ProductTexture) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.ProductTextureRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var producttexture model.ProductTexture
	if err := db.Conn.First(&producttexture, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	producttexture.Code = form.Code
	producttexture.Name = form.Name
	producttexture.Description = form.Description

	if err := db.Conn.Save(&producttexture).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateProductTextureResponse{
		ID:          producttexture.ID,
		Code:        producttexture.Code,
		Name:        producttexture.Name,
		Description: producttexture.Description,
	})
}

func (p ProductTexture) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var producttexture model.ProductTexture

	if err := db.Conn.First(&producttexture, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&producttexture, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&producttexture, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
