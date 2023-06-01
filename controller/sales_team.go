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

type SalesTeam struct{}

func (s SalesTeam) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var salesteams []model.SalesTeam

	if search != "" {
		db.Conn.Find(&salesteams, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&salesteams)
	}

	var result []dto.SalesTeamResponse
	for _, salesteam := range salesteams {
		result = append(result, dto.SalesTeamResponse{
			ID:          salesteam.ID,
			Name:        salesteam.Name,
			Description: salesteam.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (s SalesTeam) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var salesteam model.SalesTeam

	if err := db.Conn.First(&salesteam, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SalesTeamResponse{
		ID:          salesteam.ID,
		Name:        salesteam.Name,
		Description: salesteam.Description,
	})
}

func (s SalesTeam) Create(ctx *gin.Context) {

	var form dto.SalesTeamRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesteam := model.SalesTeam{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&salesteam).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSalesTeamResponse{
		ID:          salesteam.ID,
		Name:        salesteam.Name,
		Description: salesteam.Description,
	})

}

func (s SalesTeam) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SalesTeamRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var salesteam model.SalesTeam
	if err := db.Conn.First(&salesteam, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	salesteam.Name = form.Name
	salesteam.Description = form.Description

	if err := db.Conn.Save(&salesteam).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSalesTeamResponse{
		ID:          salesteam.ID,
		Name:        salesteam.Name,
		Description: salesteam.Description,
	})
}

func (s SalesTeam) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var salesteam model.SalesTeam

	if err := db.Conn.First(&salesteam, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&salesteam, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&salesteam, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
