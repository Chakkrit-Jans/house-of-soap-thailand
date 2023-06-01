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

type Sales struct{}

func (s Sales) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var sales []model.Sales

	query := db.Conn.Preload("Title").Preload("SalesTeam")

	if search != "" {
		query.Find(&sales, "code LIKE ? OR first_name LIKE ? OR last_name LIKE ? OR full_name LIKE ? OR nick_name LIKE ? OR email LIKE ? OR mobile LIKE ? OR status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		query.Find(&sales)
	}

	var result []dto.SalesResponse
	for _, sale := range sales {
		result = append(result, dto.SalesResponse{
			ID:   sale.ID,
			Code: sale.Code,
			Title: dto.TitleResponse{
				ID:          sale.Title.ID,
				Name:        sale.Title.Name,
				Description: sale.Title.Description,
			},
			FirstName: sale.FirstName,
			LastName:  sale.LastName,
			FullName:  sale.FullName,
			NickName:  sale.NickName,
			Email:     sale.Email,
			Mobile:    sale.Mobile,
			Image:     sale.Image,
			Status:    sale.Status,

			SalesTeam: dto.SalesTeamResponse{
				ID:          sale.SalesTeam.ID,
				Name:        sale.SalesTeam.Name,
				Description: sale.SalesTeam.Description,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)

}

func (s Sales) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var sales model.Sales

	if err := db.Conn.First(&sales, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SalesResponse{
		ID:   sales.ID,
		Code: sales.Code,
		Title: dto.TitleResponse{
			ID:          sales.Title.ID,
			Name:        sales.Title.Name,
			Description: sales.Title.Description,
		},
		FirstName: sales.FirstName,
		LastName:  sales.LastName,
		FullName:  sales.FullName,
		NickName:  sales.NickName,
		Email:     sales.Email,
		Mobile:    sales.Mobile,
		Image:     sales.Image,
		Status:    sales.Status,

		SalesTeam: dto.SalesTeamResponse{
			ID:          sales.SalesTeam.ID,
			Name:        sales.SalesTeam.Name,
			Description: sales.SalesTeam.Description,
		},
	})
}

func (s Sales) Create(ctx *gin.Context) {

	var form dto.SalesRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/sale/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	sale := model.Sales{
		Code:        form.Code,
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		FullName:    form.FullName,
		NickName:    form.NickName,
		Email:       form.Email,
		Mobile:      form.Mobile,
		Image:       imagePath,
		Status:      form.Status,
		TitleID:     form.TitleID,
		SalesTeamID: form.SalesTeamID,
	}

	if err := db.Conn.Create(&sale).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSalesResponse{
		ID:          sale.ID,
		Code:        sale.Code,
		TitleID:     sale.TitleID,
		FirstName:   sale.FirstName,
		LastName:    sale.LastName,
		FullName:    sale.FullName,
		NickName:    sale.NickName,
		Email:       sale.Email,
		Mobile:      sale.Mobile,
		Image:       sale.Image,
		Status:      sale.Status,
		SalesTeamID: sale.SalesTeamID,
	})

}

func (s Sales) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SalesRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sale model.Sales
	if err := db.Conn.First(&sale, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if image != nil {
		imagePath := "./uploads/sale/" + uuid.New().String()
		ctx.SaveUploadedFile(image, imagePath)
		os.Remove(sale.Image)
		sale.Image = imagePath
	}

	sale.Code = form.Code
	sale.TitleID = form.TitleID
	sale.FirstName = form.FirstName
	sale.LastName = form.LastName
	sale.FullName = form.FullName
	sale.NickName = form.NickName
	sale.Email = form.Email
	sale.Mobile = form.Mobile
	sale.Status = form.Status
	sale.SalesTeamID = form.SalesTeamID

	if err := db.Conn.Save(&sale).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSalesResponse{
		ID:          sale.ID,
		Code:        sale.Code,
		TitleID:     sale.TitleID,
		FirstName:   sale.FirstName,
		LastName:    sale.LastName,
		FullName:    sale.FullName,
		NickName:    sale.NickName,
		Email:       sale.Email,
		Mobile:      sale.Mobile,
		Image:       sale.Image,
		Status:      sale.Status,
		SalesTeamID: sale.SalesTeamID,
	})
}

func (s Sales) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var sales model.Sales

	if err := db.Conn.First(&sales, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&sales, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&sales, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
