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

type Title struct{}

func (t Title) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var titles []model.Title

	if search != "" {
		db.Conn.Find(&titles, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&titles)
	}

	var result []dto.TitleResponse
	for _, title := range titles {
		result = append(result, dto.TitleResponse{
			ID:          title.ID,
			Name:        title.Name,
			Description: title.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (t Title) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var title model.Title

	if err := db.Conn.First(&title, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.TitleResponse{
		ID:          title.ID,
		Name:        title.Name,
		Description: title.Description,
	})
}

func (t Title) Create(ctx *gin.Context) {

	var form dto.TitleRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := model.Title{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&title).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateTitleResponse{
		ID:          title.ID,
		Name:        title.Name,
		Description: title.Description,
	})

}

func (t Title) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.TitleRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var title model.Title
	if err := db.Conn.First(&title, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	title.Name = form.Name
	title.Description = form.Description

	if err := db.Conn.Save(&title).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateTitleResponse{
		ID:          title.ID,
		Name:        title.Name,
		Description: title.Description,
	})
}

func (t Title) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var title model.Title

	if err := db.Conn.First(&title, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&title, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&title, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
