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

type SystemUsers struct{}

func (u SystemUsers) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")
	var systemusers []model.SystemUsers

	query := db.Conn.Preload("Title").Preload("Sales")

	if search != "" {
		query.Find(&systemusers, "email LIKE ? OR first_name LIKE ? OR last_name LIKE ? OR full_name LIKE ? OR nick_name LIKE ? OR mobile LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		query.Find(&systemusers)
	}

	var result []dto.SystemUsersResponse
	for _, systemuser := range systemusers {
		result = append(result, dto.SystemUsersResponse{
			ID:        systemuser.ID,
			Email:     systemuser.Email,
			FirstName: systemuser.FirstName,
			LastName:  systemuser.LastName,
			FullName:  systemuser.FullName,
			NickName:  systemuser.NickName,
			Mobile:    systemuser.Mobile,
			Status:    systemuser.Status,
			Verify:    systemuser.Verify,

			Title: dto.TitleResponse{
				ID:          systemuser.Title.ID,
				Name:        systemuser.Title.Name,
				Description: systemuser.Title.Description,
			},

			Sales: dto.SalesOnlyResponse{
				ID:        systemuser.Sales.ID,
				Code:      systemuser.Sales.Code,
				FirstName: systemuser.Sales.FirstName,
				LastName:  systemuser.Sales.LastName,
				FullName:  systemuser.Sales.FullName,
				NickName:  systemuser.Sales.NickName,
				Email:     systemuser.Sales.Email,
				Mobile:    systemuser.Sales.Mobile,
				Image:     systemuser.Sales.Image,
				Status:    systemuser.Sales.Status,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (u SystemUsers) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var systemuser model.SystemUsers

	query := db.Conn.Preload("Title").Preload("Sales")

	if err := query.First(&systemuser, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.SystemUsersResponse{
		ID:        systemuser.ID,
		Email:     systemuser.Email,
		FirstName: systemuser.FirstName,
		LastName:  systemuser.LastName,
		FullName:  systemuser.FullName,
		NickName:  systemuser.NickName,
		Mobile:    systemuser.Mobile,
		Status:    systemuser.Status,
		Verify:    systemuser.Verify,

		Title: dto.TitleResponse{
			ID:          systemuser.Title.ID,
			Name:        systemuser.Title.Name,
			Description: systemuser.Title.Description,
		},

		Sales: dto.SalesOnlyResponse{
			ID:        systemuser.Sales.ID,
			Code:      systemuser.Sales.Code,
			FirstName: systemuser.Sales.FirstName,
			LastName:  systemuser.Sales.LastName,
			FullName:  systemuser.Sales.FullName,
			NickName:  systemuser.Sales.NickName,
			Email:     systemuser.Sales.Email,
			Mobile:    systemuser.Sales.Mobile,
			Image:     systemuser.Sales.Image,
			Status:    systemuser.Sales.Status,
		},
	})
}

func (u SystemUsers) Create(ctx *gin.Context) {

	var form dto.SystemUsersRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	systemuser := model.SystemUsers{
		Email:     form.Email,
		Password:  form.Password,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		FullName:  form.FullName,
		NickName:  form.NickName,
		Mobile:    form.Mobile,
		Status:    form.Status,
		Verify:    form.Verify,

		TitleID: form.TitleID,
		SalesID: form.SalesID,
	}

	if err := db.Conn.Create(&systemuser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateSystemUsersResponse{
		ID:        systemuser.ID,
		Email:     systemuser.Email,
		Password:  form.Password,
		FirstName: systemuser.FirstName,
		LastName:  systemuser.LastName,
		FullName:  systemuser.FullName,
		NickName:  systemuser.NickName,
		Mobile:    systemuser.Mobile,
		Status:    systemuser.Status,
		Verify:    systemuser.Verify,
		TitleID:   systemuser.TitleID,
		SalesID:   systemuser.SalesID,
	})

}

func (u SystemUsers) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.SystemUsersRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var systemuser model.SystemUsers
	if err := db.Conn.First(&systemuser, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	systemuser.Email = form.Email
	systemuser.FirstName = form.FirstName
	systemuser.LastName = form.LastName
	systemuser.FullName = form.FullName
	systemuser.NickName = form.NickName
	systemuser.Mobile = form.Mobile
	systemuser.Status = form.Status
	systemuser.Verify = form.Verify

	if err := db.Conn.Save(&systemuser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateSystemUsersResponse{
		ID:        systemuser.ID,
		Email:     systemuser.Email,
		FirstName: systemuser.FirstName,
		LastName:  systemuser.LastName,
		FullName:  systemuser.FullName,
		NickName:  systemuser.NickName,
		Mobile:    systemuser.Mobile,
		Status:    systemuser.Status,
		Verify:    systemuser.Verify,
	})
}

func (u SystemUsers) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var systemuser model.SystemUsers

	if err := db.Conn.First(&systemuser, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&systemuser, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&systemuser, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
