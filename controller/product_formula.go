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

type ProductFormula struct{}

func (p ProductFormula) FindAll(ctx *gin.Context) {

	productId := ctx.Query("productId")
	producttypeId := ctx.Query("producttypeId")
	producttextureId := ctx.Query("producttextureId")
	smellId := ctx.Query("smellId")
	color1Id := ctx.Query("color1Id")
	color2Id := ctx.Query("color2Id")
	color3Id := ctx.Query("color3Id")
	claimId := ctx.Query("claimId")
	howtoId := ctx.Query("howtoId")
	search := ctx.Query("search")
	var productformulas []model.ProductFormula

	query := db.Conn.Preload("Product")
	query = query.Preload("ProductType")
	query = query.Preload("ProductTexture")
	query = query.Preload("Color1")
	query = query.Preload("Color2")
	query = query.Preload("Color3")
	query = query.Preload("Smell")
	query = query.Preload("TestSize")
	query = query.Preload("Claim")
	query = query.Preload("Howto")

	// db.Joins("Company").Find(&users)
	// query = query.Joins("SmellExtracts")
	// query = query.Joins("SmellGroup")
	// query = query.Joins("SmellType")

	// .Preload("SmellExtracts").Preload("SmellGroup").Preload("SmellType")

	if productId != "" {
		query = query.Where("product_id = ?", productId)
	}

	if producttypeId != "" {
		query = query.Where("product_type_id = ?", producttypeId)
	}

	if producttextureId != "" {
		query = query.Where("product_texture_id = ?", producttextureId)
	}

	if smellId != "" {
		query = query.Where("smell_id = ?", smellId)
	}

	if color1Id != "" {
		query = query.Where("color1_id = ?", color1Id)
	}

	if color2Id != "" {
		query = query.Where("color2_id = ?", color2Id)
	}

	if color3Id != "" {
		query = query.Where("color3_id = ?", color3Id)
	}

	if claimId != "" {
		query = query.Where("claim_id = ?", claimId)
	}

	if howtoId != "" {
		query = query.Where("howto_id = ?", howtoId)
	}

	if search != "" {
		query.Find(&productformulas, "code LIKE ? OR name LIKE ? OR properties LIKE ? OR active_ingedient LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		query.Find(&productformulas)
	}

	var result []dto.ProductFormulaResponse
	for _, productformula := range productformulas {
		result = append(result, dto.ProductFormulaResponse{
			ID:              productformula.ID,
			Code:            productformula.Code,
			Name:            productformula.Name,
			Image:           productformula.Image,
			Properties:      productformula.Properties,
			ActiveIngedient: productformula.ActiveIngedient,
			UnitCost:        productformula.UnitCost,
			SalePrice:       productformula.SalePrice,

			Product: dto.ProductResponse{
				ID:    productformula.Product.ID,
				Code:  productformula.Product.Code,
				Name:  productformula.Product.Name,
				Image: productformula.Product.Image,
			},

			ProductType: dto.ProductTypeOnlyResponse{
				ID:    productformula.ProductType.ID,
				Name:  productformula.ProductType.Name,
				Image: productformula.ProductType.Image,
			},

			ProductTexture: dto.ProductTextureResponse{
				ID:          productformula.ProductTexture.ID,
				Code:        productformula.ProductTexture.Code,
				Name:        productformula.ProductTexture.Name,
				Description: productformula.ProductTexture.Description,
			},

			Smell: dto.SmellOnlyResponse{
				ID:          productformula.Smell.ID,
				Name:        productformula.Smell.Name,
				Description: productformula.Smell.Description,
			},

			Color1: dto.Color1Response{
				ID:          productformula.Color1.ID,
				Name:        productformula.Color1.Name,
				Description: productformula.Color1.Description,
			},

			Color2: dto.Color2Response{
				ID:          productformula.Color2.ID,
				Name:        productformula.Color2.Name,
				Description: productformula.Color2.Description,
			},

			Color3: dto.Color3Response{
				ID:          productformula.Color3.ID,
				Name:        productformula.Color3.Name,
				Description: productformula.Color3.Description,
			},

			TestSize: dto.TestSizeResponse{
				ID:   productformula.TestSize.ID,
				Name: productformula.TestSize.Name,
				Qty:  productformula.TestSize.Qty,
			},

			Claim: dto.ClaimResponse{
				ID:              productformula.Claim.ID,
				Code:            productformula.Claim.Code,
				Description:     productformula.Claim.Description,
				DescriptionThai: productformula.Claim.DescriptionThai,
			},

			Howto: dto.HowtoResponse{
				ID:              productformula.Howto.ID,
				Code:            productformula.Howto.Code,
				Description:     productformula.Howto.Description,
				DescriptionThai: productformula.Howto.DescriptionThai,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p ProductFormula) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var productformula model.ProductFormula

	query := db.Conn.Preload("Product").Preload("ProductType").Preload("ProductTexture")
	query.Preload("Color1").Preload("Color2").Preload("Color3")
	query.Preload("Smell").Preload("Claim").Preload("Howto")
	query.Preload("SmellExtracts").Preload("SmellGroup").Preload("SmellType")

	if err := query.First(&productformula, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ProductFormulaResponse{
		ID:              productformula.ID,
		Code:            productformula.Code,
		Name:            productformula.Name,
		Image:           productformula.Image,
		Properties:      productformula.Properties,
		ActiveIngedient: productformula.ActiveIngedient,
		UnitCost:        productformula.UnitCost,
		SalePrice:       productformula.SalePrice,

		Product: dto.ProductResponse{
			ID:    productformula.Product.ID,
			Code:  productformula.Product.Code,
			Name:  productformula.Product.Name,
			Image: productformula.Product.Image,
		},
		ProductType: dto.ProductTypeOnlyResponse{
			ID:    productformula.ProductType.ID,
			Name:  productformula.ProductType.Name,
			Image: productformula.ProductType.Image,
		},
		ProductTexture: dto.ProductTextureResponse{
			ID:          productformula.ProductTexture.ID,
			Code:        productformula.ProductTexture.Code,
			Name:        productformula.ProductTexture.Name,
			Description: productformula.ProductTexture.Description,
		},
		Smell: dto.SmellOnlyResponse{
			ID:          productformula.Smell.ID,
			Name:        productformula.Smell.Name,
			Description: productformula.Smell.Description,
		},

		Color1: dto.Color1Response{
			ID:          productformula.Color1.ID,
			Name:        productformula.Color1.Name,
			Description: productformula.Color1.Description,
		},
		Color2: dto.Color2Response{
			ID:          productformula.Color2.ID,
			Name:        productformula.Color2.Name,
			Description: productformula.Color2.Description,
		},
		Color3: dto.Color3Response{
			ID:          productformula.Color3.ID,
			Name:        productformula.Color3.Name,
			Description: productformula.Color3.Description,
		},
		Claim: dto.ClaimResponse{
			ID:              productformula.Claim.ID,
			Code:            productformula.Claim.Code,
			Description:     productformula.Claim.Description,
			DescriptionThai: productformula.Claim.DescriptionThai,
		},
		Howto: dto.HowtoResponse{
			ID:              productformula.Howto.ID,
			Code:            productformula.Howto.Code,
			Description:     productformula.Howto.Description,
			DescriptionThai: productformula.Howto.DescriptionThai,
		},
	})
}

func (p ProductFormula) Create(ctx *gin.Context) {

	var form dto.ProductFormulaRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/productsformula/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	productformula := model.ProductFormula{
		Code:            form.Code,
		Name:            form.Name,
		Image:           imagePath,
		Properties:      form.Properties,
		ActiveIngedient: form.Properties,
		UnitCost:        form.UnitCost,
		SalePrice:       form.SalePrice,

		ProductID:        form.ProductID,
		ProductTypeID:    form.ProductTypeID,
		ProductTextureID: form.ProductTextureID,
		SmellID:          form.SmellID,
		Color1ID:         form.Color1ID,
		Color2ID:         form.Color2ID,
		Color3ID:         form.Color3ID,
		ClaimID:          form.ClaimID,
		HowtoID:          form.HowtoID,
	}

	if err := db.Conn.Create(&productformula).Error; err != nil {
		os.Remove(productformula.Image)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateProductFormulaResponse{
		ID:              productformula.ID,
		Code:            productformula.Code,
		Name:            productformula.Name,
		Image:           productformula.Image,
		Properties:      productformula.Properties,
		ActiveIngedient: productformula.ActiveIngedient,
		UnitCost:        productformula.UnitCost,
		SalePrice:       productformula.SalePrice,

		ProductID:        productformula.ProductID,
		ProductTypeID:    productformula.ProductTypeID,
		ProductTextureID: productformula.ProductTextureID,
		SmellID:          productformula.SmellID,
		Color1ID:         productformula.Color1ID,
		Color2ID:         productformula.Color2ID,
		Color3ID:         productformula.Color3ID,
		ClaimID:          productformula.ClaimID,
		HowtoID:          productformula.HowtoID,
	})

}

func (p ProductFormula) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.ProductFormulaRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productformula model.ProductFormula
	if err := db.Conn.First(&productformula, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if image != nil {
		imagePath := "./uploads/productsformula/" + uuid.New().String()
		ctx.SaveUploadedFile(image, imagePath)
		os.Remove(productformula.Image)
		productformula.Image = imagePath
	}

	productformula.Code = form.Code
	productformula.Name = form.Name
	productformula.Properties = form.Properties
	productformula.ActiveIngedient = form.Properties
	productformula.UnitCost = form.UnitCost
	productformula.SalePrice = form.SalePrice

	productformula.ProductID = form.ProductID
	productformula.ProductTypeID = form.ProductTypeID
	productformula.ProductTextureID = form.ProductTextureID
	productformula.SmellID = form.SmellID
	productformula.Color1ID = form.Color1ID
	productformula.Color2ID = form.Color2ID
	productformula.Color3ID = form.Color3ID
	productformula.ClaimID = form.ClaimID
	productformula.HowtoID = form.HowtoID

	if err := db.Conn.Save(&productformula).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateProductFormulaResponse{
		ID:              productformula.ID,
		Code:            productformula.Code,
		Name:            productformula.Name,
		Image:           productformula.Image,
		Properties:      productformula.Properties,
		ActiveIngedient: productformula.ActiveIngedient,
		UnitCost:        productformula.UnitCost,
		SalePrice:       productformula.SalePrice,

		ProductID:        productformula.ProductID,
		ProductTypeID:    productformula.ProductTypeID,
		ProductTextureID: productformula.ProductTextureID,
		SmellID:          productformula.SmellID,
		Color1ID:         productformula.Color1ID,
		Color2ID:         productformula.Color2ID,
		Color3ID:         productformula.Color3ID,
		ClaimID:          productformula.ClaimID,
		HowtoID:          productformula.HowtoID,
	})
}

func (p ProductFormula) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var productformula model.ProductFormula

	if err := db.Conn.First(&productformula, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&productformula, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&product, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
