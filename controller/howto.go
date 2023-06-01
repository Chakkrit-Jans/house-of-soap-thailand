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

type Howto struct{}

func (c Howto) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var howtos []model.Howto

	if search != "" {
		db.Conn.Find(&howtos, "code LIKE ? OR description LIKE ? OR description_thai LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&howtos)
	}

	var result []dto.HowtoResponse
	for _, howto := range howtos {
		result = append(result, dto.HowtoResponse{
			ID:              howto.ID,
			Code:            howto.Code,
			Description:     howto.Description,
			DescriptionThai: howto.DescriptionThai,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Howto) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var howto model.Howto

	if err := db.Conn.First(&howto, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.HowtoResponse{
		ID:              howto.ID,
		Code:            howto.Code,
		Description:     howto.Description,
		DescriptionThai: howto.DescriptionThai,
	})
}

func (c Howto) Create(ctx *gin.Context) {

	var form dto.HowtoRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	howto := model.Howto{
		Code:            form.Code,
		Description:     form.Description,
		DescriptionThai: form.DescriptionThai,
	}

	if err := db.Conn.Create(&howto).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateHowtoResponse{
		ID:              howto.ID,
		Code:            howto.Code,
		Description:     howto.Description,
		DescriptionThai: howto.DescriptionThai,
	})

}

func (c Howto) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.HowtoRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var howto model.Howto
	if err := db.Conn.First(&howto, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	howto.Code = form.Code
	howto.Description = form.Description
	howto.DescriptionThai = form.DescriptionThai

	if err := db.Conn.Save(&howto).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateHowtoResponse{
		ID:              howto.ID,
		Code:            howto.Code,
		Description:     howto.Description,
		DescriptionThai: howto.DescriptionThai,
	})
}

func (c Howto) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var howto model.Howto

	if err := db.Conn.First(&howto, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&howto, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&howto, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
