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

type TestSize struct{}

func (t TestSize) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var testsizes []model.TestSize

	query := db.Conn.Preload("Units")

	if search != "" {
		query.Find(&testsizes, "name LIKE ? ", "%"+search+"%")
	} else {
		query.Find(&testsizes)
	}

	var result []dto.TestSizeResponse
	for _, testsize := range testsizes {
		result = append(result, dto.TestSizeResponse{
			ID:   testsize.ID,
			Name: testsize.Name,
			Qty:  testsize.Qty,

			Units: dto.UnitResponse{
				ID:          testsize.Units.ID,
				Name:        testsize.Units.Name,
				Description: testsize.Units.Description,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (t TestSize) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var testsize model.TestSize

	if err := db.Conn.First(&testsize, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.TestSizeResponse{
		ID:   testsize.ID,
		Name: testsize.Name,
		Qty:  testsize.Qty,

		Units: dto.UnitResponse{
			ID:          testsize.Units.ID,
			Name:        testsize.Units.Name,
			Description: testsize.Units.Description,
		},
	})
}

func (t TestSize) Create(ctx *gin.Context) {

	var form dto.TestSizeRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testsize := model.TestSize{
		Name:    form.Name,
		Qty:     form.Qty,
		UnitsID: form.UnitsID,
	}

	if err := db.Conn.Create(&testsize).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateTestSizeResponse{
		ID:      testsize.ID,
		Name:    testsize.Name,
		Qty:     testsize.Qty,
		UnitsID: testsize.UnitsID,
	})

}

func (t TestSize) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.TestSizeRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var testsize model.TestSize
	if err := db.Conn.First(&testsize, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	testsize.Name = form.Name
	testsize.Qty = form.Qty
	testsize.UnitsID = form.UnitsID

	if err := db.Conn.Save(&testsize).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateTestSizeResponse{
		ID:      testsize.ID,
		Name:    testsize.Name,
		Qty:     testsize.Qty,
		UnitsID: testsize.UnitsID,
	})
}

func (t TestSize) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var testsize model.TestSize

	if err := db.Conn.First(&testsize, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&testsize, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&testsize, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
