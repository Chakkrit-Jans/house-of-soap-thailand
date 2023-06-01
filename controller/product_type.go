package controller

import (
	"errors"
	"hst-api/db"
	"hst-api/dto"
	"hst-api/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductType struct{}

func (p ProductType) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var producttypes []model.ProductType

	query := db.Conn.Preload("Product")

	if search != "" {
		query.Find(&producttypes, "name LIKE ?", "%"+search+"%")
	} else {
		query.Find(&producttypes)
	}

	var result []dto.ProductTypeResponse
	for _, producttype := range producttypes {
		result = append(result, dto.ProductTypeResponse{
			ID:    producttype.ID,
			Name:  producttype.Name,
			Image: producttype.Image,

			Product: dto.ProductResponse{
				ID:   producttype.Product.ID,
				Code: producttype.Product.Code,
				Name: producttype.Product.Name,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p ProductType) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var producttype model.ProductType

	query := db.Conn.Preload("Product")

	if err := query.First(&producttype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductTypeResponse{
		ID:    producttype.ID,
		Name:  producttype.Name,
		Image: producttype.Image,

		Product: dto.ProductResponse{
			ID:   producttype.Product.ID,
			Code: producttype.Product.Code,
			Name: producttype.Product.Name,
		},
	})
}

func (p ProductType) Create(ctx *gin.Context) {

	var form dto.ProductTypeRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/producttype/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	producttype := model.ProductType{
		Name:  form.Name,
		Image: imagePath,

		ProductID: form.ProductID,
	}

	if err := db.Conn.Create(&producttype).Error; err != nil {
		os.Remove(producttype.Image)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateProductTypeResponse{
		ID:        producttype.ID,
		Name:      producttype.Name,
		Image:     producttype.Image,
		ProductID: producttype.ProductID,
	})

}

func (p ProductType) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.ProductTypeRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var producttype model.ProductType
	if err := db.Conn.First(&producttype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if image != nil {
		imagePath := "./uploads/producttype/" + uuid.New().String()
		ctx.SaveUploadedFile(image, imagePath)
		os.Remove(producttype.Image)
		producttype.Image = imagePath
	}

	producttype.Name = form.Name
	producttype.ProductID = form.ProductID

	if err := db.Conn.Save(&producttype).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateProductTypeResponse{
		ID:        producttype.ID,
		Name:      producttype.Name,
		Image:     producttype.Image,
		ProductID: producttype.ProductID,
	})
}

func (p ProductType) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var producttype model.ProductType

	if err := db.Conn.First(&producttype, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&producttype, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&producttype, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
