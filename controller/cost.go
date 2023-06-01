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

type Cost struct{}

func (c Cost) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var costs []model.Cost
	query := db.Conn

	if search != "" {
		query.Find(&costs, "name LIKE ?", "%"+search+"%")
	} else {
		query.Find(&costs)
	}

	var result []dto.CostResponse
	for _, cost := range costs {
		result = append(result, dto.CostResponse{
			ID:        cost.ID,
			Code:      cost.Code,
			Name:      cost.Name,
			CostType:  cost.CostType,
			CostGroup: cost.CostGroup,
			Amount:    cost.Amount,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Cost) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var cost model.Cost

	if err := db.Conn.First(&cost, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CostResponse{
		ID:        cost.ID,
		Code:      cost.Code,
		Name:      cost.Name,
		CostType:  cost.CostType,
		CostGroup: cost.CostGroup,
		Amount:    cost.Amount,
	})
}

func (c Cost) Create(ctx *gin.Context) {

	var form dto.CostRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cost := model.Cost{
		Code:      form.Code,
		Name:      form.Name,
		CostType:  form.CostType,
		CostGroup: form.CostGroup,
		Amount:    form.Amount,
	}

	if err := db.Conn.Create(&cost).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateCostResponse{
		ID:        cost.ID,
		Code:      cost.Code,
		Name:      cost.Name,
		CostType:  cost.CostType,
		CostGroup: cost.CostGroup,
		Amount:    cost.Amount,
	})

}

func (c Cost) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.CostRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cost model.Cost
	if err := db.Conn.First(&cost, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	cost.Code = form.Code
	cost.Name = form.Name
	cost.CostGroup = form.CostGroup
	cost.CostType = form.CostType
	cost.Amount = form.Amount

	if err := db.Conn.Save(&cost).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateCostResponse{
		ID:        cost.ID,
		Code:      cost.Code,
		Name:      cost.Name,
		CostGroup: cost.CostGroup,
		CostType:  cost.CostType,
		Amount:    cost.Amount,
	})
}

func (c Cost) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var cost model.Cost

	if err := db.Conn.First(&cost, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&cost, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&cost, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
