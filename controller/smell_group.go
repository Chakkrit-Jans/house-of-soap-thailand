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

type SmellGroup struct{}

func (s SmellGroup) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var smellgroups []model.SmellGroup

	if search != "" {
		db.Conn.Find(&smellgroups, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&smellgroups)
	}

	var result []dto.SmellGroupResponse
	for _, smellgroup := range smellgroups {
		result = append(result, dto.SmellGroupResponse{
			ID:          smellgroup.ID,
			Name:        smellgroup.Name,
			Description: smellgroup.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (s SmellGroup) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var smellgroup model.SmellGroup

	if err := db.Conn.First(&smellgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SmellGroupResponse{
		ID:          smellgroup.ID,
		Name:        smellgroup.Name,
		Description: smellgroup.Description,
	})
}

func (s SmellGroup) Create(ctx *gin.Context) {

	var form dto.SmellGroupRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	smellgroup := model.SmellGroup{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&smellgroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSmellGroupResponse{
		ID:          smellgroup.ID,
		Name:        smellgroup.Name,
		Description: smellgroup.Description,
	})
}

func (s SmellGroup) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SmellGroupRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var smellgroup model.SmellGroup
	if err := db.Conn.First(&smellgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	smellgroup.Name = form.Name
	smellgroup.Description = form.Description

	if err := db.Conn.Save(&smellgroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSmellGroupResponse{
		ID:          smellgroup.ID,
		Name:        smellgroup.Name,
		Description: smellgroup.Description,
	})
}

func (s SmellGroup) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var smellgroup model.SmellGroup

	if err := db.Conn.First(&smellgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&smellgroup, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&smellgroup, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
