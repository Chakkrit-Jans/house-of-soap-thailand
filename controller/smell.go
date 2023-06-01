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

type Smell struct{}

func (s Smell) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var smells []model.Smell

	query := db.Conn.Preload("SmellExtracts").Preload("SmellType").Preload("SmellGroup")

	if search != "" {
		query.Find(&smells, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		query.Find(&smells)
	}

	var result []dto.SmellResponse
	for _, smell := range smells {
		result = append(result, dto.SmellResponse{
			ID:          smell.ID,
			Name:        smell.Name,
			Description: smell.Description,

			SmellExtracts: dto.SmellExtractsResponse{
				ID:          smell.SmellExtracts.ID,
				Name:        smell.SmellExtracts.Name,
				Description: smell.SmellExtracts.Description,
			},
			SmellType: dto.SmellTypeResponse{
				ID:          smell.SmellType.ID,
				Name:        smell.SmellType.Name,
				Description: smell.SmellType.Description,
			},
			SmellGroup: dto.SmellGroupResponse{
				ID:          smell.SmellGroup.ID,
				Name:        smell.SmellGroup.Name,
				Description: smell.SmellGroup.Description,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (s Smell) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var smell model.Smell

	query := db.Conn.Preload("SmellExtracts").Preload("SmellType").Preload("SmellGroup")

	if err := query.First(&smell, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SmellResponse{
		ID:          smell.ID,
		Name:        smell.Name,
		Description: smell.Description,

		SmellExtracts: dto.SmellExtractsResponse{
			ID:          smell.SmellExtracts.ID,
			Name:        smell.SmellExtracts.Name,
			Description: smell.SmellExtracts.Description,
		},
		SmellType: dto.SmellTypeResponse{
			ID:          smell.SmellType.ID,
			Name:        smell.SmellType.Name,
			Description: smell.SmellType.Description,
		},
		SmellGroup: dto.SmellGroupResponse{
			ID:          smell.SmellGroup.ID,
			Name:        smell.SmellGroup.Name,
			Description: smell.SmellGroup.Description,
		},
	})
}

func (s Smell) Create(ctx *gin.Context) {

	var form dto.SmellRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	smell := model.Smell{
		Name:            form.Name,
		Description:     form.Description,
		SmellExtractsID: form.SmellExtractsID,
		SmellTypeID:     form.SmellTypeID,
		SmellGroupID:    form.SmellGroupID,
	}

	if err := db.Conn.Create(&smell).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSmellResponse{
		ID:              smell.ID,
		Name:            smell.Name,
		Description:     smell.Description,
		SmellExtractsID: smell.SmellExtractsID,
		SmellTypeID:     smell.SmellTypeID,
		SmellGroupID:    smell.SmellGroupID,
	})

}

func (s Smell) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SmellRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var smell model.Smell
	if err := db.Conn.First(&smell, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	smell.Name = form.Name
	smell.Description = form.Description
	smell.SmellExtractsID = form.SmellExtractsID
	smell.SmellTypeID = form.SmellTypeID
	smell.SmellGroupID = form.SmellGroupID

	if err := db.Conn.Save(&smell).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSmellResponse{
		ID:              smell.ID,
		Name:            smell.Name,
		Description:     smell.Description,
		SmellExtractsID: smell.SmellExtractsID,
		SmellTypeID:     smell.SmellTypeID,
		SmellGroupID:    smell.SmellGroupID,
	})
}

func (s Smell) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var smell model.Smell

	if err := db.Conn.First(&smell, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&smell, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&smell, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
