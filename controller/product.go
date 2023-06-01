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

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var products []model.Product

	if search != "" {
		db.Conn.Find(&products, "code LIKE ? OR name LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&products)
	}

	var result []dto.ProductResponse
	for _, product := range products {
		result = append(result, dto.ProductResponse{
			ID:    product.ID,
			Code:  product.Code,
			Name:  product.Name,
			Image: product.Image,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p Product) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var product model.Product

	if err := db.Conn.First(&product, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductResponse{
		ID:    product.ID,
		Code:  product.Code,
		Name:  product.Name,
		Image: product.Image,
	})
}

func (p Product) Create(ctx *gin.Context) {

	var form dto.ProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/products/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	product := model.Product{
		Code:  form.Code,
		Name:  form.Name,
		Image: imagePath,
	}

	if err := db.Conn.Create(&product).Error; err != nil {
		os.Remove(product.Image)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateProductResponse{
		Code:  form.Code,
		Name:  form.Name,
		Image: imagePath,
	})

}

func (p Product) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.ProductRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product model.Product
	if err := db.Conn.First(&product, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if image != nil {
		imagePath := "./uploads/products/" + uuid.New().String()
		ctx.SaveUploadedFile(image, imagePath)
		os.Remove(product.Image)
		product.Image = imagePath
	}

	product.Code = form.Code
	product.Name = form.Name

	if err := db.Conn.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateProductResponse{
		ID:    product.ID,
		Code:  product.Code,
		Name:  product.Name,
		Image: product.Image,
	})
}

func (p Product) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var product model.Product

	if err := db.Conn.First(&product, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&product, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&product, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
