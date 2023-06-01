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

type Claims struct{}

func (c Claims) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var claims []model.Claim

	if search != "" {
		db.Conn.Find(&claims, "code LIKE ? OR description LIKE ? OR description_thai LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&claims)
	}

	var result []dto.ClaimResponse
	for _, claim := range claims {
		result = append(result, dto.ClaimResponse{
			ID:              claim.ID,
			Code:            claim.Code,
			Description:     claim.Description,
			DescriptionThai: claim.DescriptionThai,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Claims) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var claim model.Claim

	if err := db.Conn.First(&claim, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ClaimResponse{
		ID:              claim.ID,
		Code:            claim.Code,
		Description:     claim.Description,
		DescriptionThai: claim.DescriptionThai,
	})
}

func (c Claims) Create(ctx *gin.Context) {

	var form dto.ClaimRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claim := model.Claim{
		Code:            form.Code,
		Description:     form.Description,
		DescriptionThai: form.DescriptionThai,
	}

	if err := db.Conn.Create(&claim).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateClaimResponse{
		ID:              claim.ID,
		Code:            claim.Code,
		Description:     claim.Description,
		DescriptionThai: claim.DescriptionThai,
	})

}

func (c Claims) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.ClaimRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var claim model.Claim
	if err := db.Conn.First(&claim, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	claim.Code = form.Code
	claim.Description = form.Description
	claim.DescriptionThai = form.DescriptionThai

	if err := db.Conn.Save(&claim).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateClaimResponse{
		ID:              claim.ID,
		Code:            claim.Code,
		Description:     claim.Description,
		DescriptionThai: claim.DescriptionThai,
	})
}

func (c Claims) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var claim model.Claim

	if err := db.Conn.First(&claim, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&claim, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&claim, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
