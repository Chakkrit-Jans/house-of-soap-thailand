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

type UsersGroup struct{}

func (u UsersGroup) FindAll(ctx *gin.Context) {

	search := ctx.Query("search")

	var usersgroups []model.UsersGroup

	if search != "" {
		db.Conn.Find(&usersgroups, "name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	} else {
		db.Conn.Find(&usersgroups)
	}

	var result []dto.UsersGroupResponse
	for _, usersgroup := range usersgroups {
		result = append(result, dto.UsersGroupResponse{
			ID:          usersgroup.ID,
			Name:        usersgroup.Name,
			Description: usersgroup.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (u UsersGroup) FindOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var usersgroup model.UsersGroup

	if err := db.Conn.First(&usersgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.UsersGroupResponse{
		ID:          usersgroup.ID,
		Name:        usersgroup.Name,
		Description: usersgroup.Description,
	})
}

func (u UsersGroup) Create(ctx *gin.Context) {

	var form dto.UsersGroupRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersgroup := model.UsersGroup{
		Name:        form.Name,
		Description: form.Description,
	}

	if err := db.Conn.Create(&usersgroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateUsersGroupResponse{
		ID:          usersgroup.ID,
		Name:        usersgroup.Name,
		Description: usersgroup.Description,
	})

}

func (u UsersGroup) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var form dto.UsersGroupRequest

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usersgroup model.UsersGroup
	if err := db.Conn.First(&usersgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	usersgroup.Name = form.Name
	usersgroup.Description = form.Description

	if err := db.Conn.Save(&usersgroup).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateUsersGroupResponse{
		ID:          usersgroup.ID,
		Name:        usersgroup.Name,
		Description: usersgroup.Description,
	})
}

func (u UsersGroup) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	var usersgroup model.UsersGroup

	if err := db.Conn.First(&usersgroup, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Conn.Unscoped().Delete(&usersgroup, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// db.Conn.Unscoped().Delete(&usersgroup, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
