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

type Customers struct{}

func (c Customers) FindAll(ctx *gin.Context) {

	TitleId := ctx.Query("TitleId")
	CustomerGroupId := ctx.Query("CustomerGroupId")
	CustomerTypeId := ctx.Query("CustomerTypeId")
	vatId := ctx.Query("vatId")
	search := ctx.Query("search")

	var customers []model.Customers
	query := db.Conn.Preload("Title").Preload("CustomerGroup").Preload("CustomerType").Preload("Vat")

	if TitleId != "" {
		query = query.Where("title_id = ?", TitleId)
	}

	if CustomerGroupId != "" {
		query = query.Where("customer_group_id = ?", CustomerGroupId)
	}

	if CustomerTypeId != "" {
		query = query.Where("customer_type_id = ?", CustomerTypeId)
	}

	if vatId != "" {
		query = query.Where("vat_id = ?", vatId)
	}

	if search != "" {
		query = query.Where("code LIKE ? OR first_name LIKE ? OR last_name LIKE ? OR full_name LIKE ? OR nick_name LIKE ? OR email LIKE ? OR mobile LIKE ? OR status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query.Find(&customers)

	var result []dto.CustomerResponse
	for _, customer := range customers {
		result = append(result, dto.CustomerResponse{
			ID:   customer.ID,
			Code: customer.Code,

			Title: dto.TitleResponse{
				ID:   customer.Title.ID,
				Name: customer.Title.Name,
			},
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			FullName:  customer.FullName,
			NickName:  customer.NickName,
			Email:     customer.Email,
			Mobile:    customer.Mobile,
			Status:    customer.Status,

			CustomerGroup: dto.CustomerGroupOnlyResponse{
				ID:          customer.CustomerGroup.ID,
				Name:        customer.CustomerGroup.Name,
				Description: customer.CustomerGroup.Description,
			},

			CustomerType: dto.CustomerTypeResponse{
				ID:          customer.CustomerType.ID,
				Name:        customer.CustomerType.Name,
				Description: customer.CustomerType.Description,
			},

			Vat: dto.VatResponse{
				ID:         customer.Vat.ID,
				Name:       customer.Vat.Name,
				Percentage: customer.Vat.Percentage,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)

}

func (c Customers) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var customer model.Customers
	query := db.Conn.Preload("Title").Preload("CustomerGroup").Preload("CustomerType").Preload("Vat")

	if err := query.First(&customer, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CustomerResponse{
		ID:   customer.ID,
		Code: customer.Code,
		Title: dto.TitleResponse{
			ID:   customer.Title.ID,
			Name: customer.Title.Name,
		},
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		FullName:  customer.FullName,
		NickName:  customer.NickName,
		Email:     customer.Email,
		Mobile:    customer.Mobile,
		Status:    customer.Status,

		CustomerGroup: dto.CustomerGroupOnlyResponse{
			ID:          customer.CustomerGroup.ID,
			Name:        customer.CustomerGroup.Name,
			Description: customer.CustomerGroup.Description,
		},

		CustomerType: dto.CustomerTypeResponse{
			ID:          customer.CustomerType.ID,
			Name:        customer.CustomerType.Name,
			Description: customer.CustomerType.Description,
		},

		Vat: dto.VatResponse{
			ID:         customer.Vat.ID,
			Name:       customer.Vat.Name,
			Percentage: customer.Vat.Percentage,
		},
	})

}

func (c Customers) Create(ctx *gin.Context) {

	var form dto.CustomerRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := model.Customers{
		Code:      form.Code,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		FullName:  form.FullName,
		NickName:  form.NickName,
		Email:     form.Email,
		Mobile:    form.Mobile,

		TitleID:         form.TitleID,
		CustomerGroupID: form.CustomerGroupID,
		CustomerTypeID:  form.CustomerTypeID,
		VatID:           form.VatID,
	}

	if err := db.Conn.Create(&customer).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateCustomerResponse{
		ID:        customer.ID,
		Code:      customer.Code,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		FullName:  customer.FullName,
		NickName:  customer.NickName,
		Email:     customer.Email,
		Mobile:    customer.Mobile,

		TitleID:         customer.TitleID,
		CustomerGroupID: customer.CustomerGroupID,
		CustomerTypeID:  customer.CustomerTypeID,
		VatID:           customer.VatID,
	})

}

func (c Customers) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.CustomerRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer model.Customers

	if err := db.Conn.First(&customer, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	customer.Code = form.Code
	customer.FirstName = form.FirstName
	customer.LastName = form.LastName
	customer.FullName = form.FullName
	customer.NickName = form.NickName
	customer.Email = form.Email
	customer.Mobile = form.Mobile

	customer.TitleID = form.TitleID
	customer.CustomerGroupID = form.CustomerGroupID
	customer.CustomerTypeID = form.CustomerTypeID
	customer.VatID = form.VatID

	if err := db.Conn.Save(&customer).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateCustomerResponse{
		ID:        customer.ID,
		Code:      customer.Code,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		FullName:  customer.FullName,
		NickName:  customer.NickName,
		Email:     customer.Email,
		Mobile:    customer.Mobile,

		TitleID:         customer.TitleID,
		CustomerGroupID: customer.CustomerGroupID,
		CustomerTypeID:  customer.CustomerTypeID,
		VatID:           customer.VatID,
	})
}

func (c Customers) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var customers model.Customers

	if err := db.Conn.First(&customers, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&customers, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&customers, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
