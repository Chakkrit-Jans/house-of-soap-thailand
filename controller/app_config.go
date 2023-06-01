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

type AppConfig struct{}

func (a AppConfig) FindAll(ctx *gin.Context) {

	var appconfigs []model.AppConfig

	db.Conn.Find(&appconfigs)

	var result []dto.AppConfigResponse
	for _, appconfig := range appconfigs {
		result = append(result, dto.AppConfigResponse{
			ID:                    appconfig.ID,
			ProductionCostPerItem: appconfig.ProductionCostPerItem,
			ProductionCostPerUnit: appconfig.ProductionCostPerUnit,
			PercentageProfit:      appconfig.PercentageProfit,
		})
	}

	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (a AppConfig) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var appconfig model.AppConfig

	if err := db.Conn.First(&appconfig, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.AppConfigResponse{
		ID:                    appconfig.ID,
		ProductionCostPerItem: appconfig.ProductionCostPerItem,
		ProductionCostPerUnit: appconfig.ProductionCostPerUnit,
		PercentageProfit:      appconfig.PercentageProfit,
	})
}

func (a AppConfig) Create(ctx *gin.Context) {

	var form dto.AppConfigRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appconfig := model.AppConfig{
		ProductionCostPerItem: form.ProductionCostPerItem,
		ProductionCostPerUnit: form.ProductionCostPerUnit,
		PercentageProfit:      form.PercentageProfit,
	}

	if err := db.Conn.Create(&appconfig).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateAppConfigResponse{
		ID:                    appconfig.ID,
		ProductionCostPerItem: appconfig.ProductionCostPerItem,
		ProductionCostPerUnit: appconfig.ProductionCostPerUnit,
		PercentageProfit:      appconfig.PercentageProfit,
	})

}

func (a AppConfig) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.AppConfigRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var appconfig model.AppConfig
	if err := db.Conn.First(&appconfig, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	appconfig.ProductionCostPerItem = form.ProductionCostPerItem
	appconfig.ProductionCostPerUnit = form.ProductionCostPerUnit
	appconfig.PercentageProfit = form.PercentageProfit

	if err := db.Conn.Save(&appconfig).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateAppConfigResponse{
		ID:                    appconfig.ID,
		ProductionCostPerItem: appconfig.ProductionCostPerItem,
		ProductionCostPerUnit: appconfig.ProductionCostPerUnit,
		PercentageProfit:      appconfig.PercentageProfit,
	})
}

func (a AppConfig) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var appconfig model.AppConfig

	if err := db.Conn.First(&appconfig, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&appconfig, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&appconfig, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
