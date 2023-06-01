package main

import (
	"hst-api/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {

	//========== App config ==========
	appconfigController := controller.AppConfig{}
	appconfigGroup := r.Group("/appconfig")
	appconfigGroup.GET("", appconfigController.FindAll)
	appconfigGroup.GET("/:id", appconfigController.FindOne)
	appconfigGroup.POST("", appconfigController.Create)
	appconfigGroup.PATCH("/:id", appconfigController.Update)
	appconfigGroup.DELETE("/:id", appconfigController.Delete)

	//========== Claims ==========
	claimController := controller.Claims{}
	claimGroup := r.Group("/claims")
	claimGroup.GET("", claimController.FindAll)
	claimGroup.GET("/:id", claimController.FindOne)
	claimGroup.POST("", claimController.Create)
	claimGroup.PATCH("/:id", claimController.Update)
	claimGroup.DELETE("/:id", claimController.Delete)

	//========== Clolor ==========
	color1Controller := controller.Color1{}
	color1Group := r.Group("/color1")
	color1Group.GET("", color1Controller.FindAll)
	color1Group.GET("/:id", color1Controller.FindOne)
	color1Group.POST("", color1Controller.Create)
	color1Group.PATCH("/:id", color1Controller.Update)
	color1Group.DELETE("/:id", color1Controller.Delete)

	color2Controller := controller.Color2{}
	color2Group := r.Group("/color2")
	color2Group.GET("", color2Controller.FindAll)
	color2Group.GET("/:id", color2Controller.FindOne)
	color2Group.POST("", color2Controller.Create)
	color2Group.PATCH("/:id", color2Controller.Update)
	color2Group.DELETE("/:id", color2Controller.Delete)

	color3Controller := controller.Color3{}
	color3Group := r.Group("/color3")
	color3Group.GET("", color3Controller.FindAll)
	color3Group.GET("/:id", color3Controller.FindOne)
	color3Group.POST("", color3Controller.Create)
	color3Group.PATCH("/:id", color3Controller.Update)
	color3Group.DELETE("/:id", color3Controller.Delete)

	//========== Cost ==========
	costController := controller.Cost{}
	costGroup := r.Group("/cost")
	costGroup.GET("", costController.FindAll)
	costGroup.GET("/:id", costController.FindOne)
	costGroup.POST("", costController.Create)
	costGroup.PATCH("/:id", costController.Update)
	costGroup.DELETE("/:id", costController.Delete)

	//========== Vat ==========
	vatController := controller.Vat{}
	vatGroup := r.Group("/vat")
	vatGroup.GET("", vatController.FindAll)
	vatGroup.GET("/:id", vatController.FindOne)
	vatGroup.POST("", vatController.Create)
	vatGroup.PATCH("/:id", vatController.Update)
	vatGroup.DELETE("/:id", vatController.Delete)

	//========== Discount ==========
	discountController := controller.Discount{}
	discountGroup := r.Group("/discount")
	discountGroup.GET("", discountController.FindAll)
	discountGroup.GET("/:id", discountController.FindOne)
	discountGroup.POST("", discountController.Create)
	discountGroup.PATCH("/:id", discountController.Update)
	discountGroup.DELETE("/:id", discountController.Delete)

	//========== Howto ==========
	howtoController := controller.Howto{}
	howtoGroup := r.Group("/howto")
	howtoGroup.GET("", howtoController.FindAll)
	howtoGroup.GET("/:id", howtoController.FindOne)
	howtoGroup.POST("", howtoController.Create)
	howtoGroup.PATCH("/:id", howtoController.Update)
	howtoGroup.DELETE("/:id", howtoController.Delete)

	//========== Title ==========
	titleController := controller.Title{}
	titleGroup := r.Group("/title")
	titleGroup.GET("", titleController.FindAll)
	titleGroup.GET("/:id", titleController.FindOne)
	titleGroup.POST("", titleController.Create)
	titleGroup.PATCH("/:id", titleController.Update)
	titleGroup.DELETE("/:id", titleController.Delete)

	//========== Customer ==========
	custgroupController := controller.CustomerGroup{}
	custgroupGroup := r.Group("/custgroup")
	custgroupGroup.GET("", custgroupController.FindAll)
	custgroupGroup.GET("/:id", custgroupController.FindOne)
	custgroupGroup.POST("", custgroupController.Create)
	custgroupGroup.PATCH("/:id", custgroupController.Update)
	custgroupGroup.DELETE("/:id", custgroupController.Delete)

	custtypeController := controller.CustomerType{}
	custtypeGroup := r.Group("/custtype")
	custtypeGroup.GET("", custtypeController.FindAll)
	custtypeGroup.GET("/:id", custtypeController.FindOne)
	custtypeGroup.POST("", custtypeController.Create)
	custtypeGroup.PATCH("/:id", custtypeController.Update)
	custtypeGroup.DELETE("/:id", custtypeController.Delete)

	customerController := controller.Customers{}
	customerGroup := r.Group("/customers")
	customerGroup.GET("", customerController.FindAll)
	customerGroup.GET("/:id", customerController.FindOne)
	customerGroup.POST("", customerController.Create)
	customerGroup.PATCH("/:id", customerController.Update)
	customerGroup.DELETE("/:id", customerController.Delete)

	//========== Product ==========
	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)

	//========== Product Type ==========
	producttypeController := controller.ProductType{}
	producttypeGroup := r.Group("/producttype")
	producttypeGroup.GET("", producttypeController.FindAll)
	producttypeGroup.GET("/:id", producttypeController.FindOne)
	producttypeGroup.POST("", producttypeController.Create)
	producttypeGroup.PATCH("/:id", producttypeController.Update)
	producttypeGroup.DELETE("/:id", producttypeController.Delete)

	//========== Product Texture ==========
	producttextureController := controller.ProductTexture{}
	producttextureGroup := r.Group("/producttexture")
	producttextureGroup.GET("", producttextureController.FindAll)
	producttextureGroup.GET("/:id", producttextureController.FindOne)
	producttextureGroup.POST("", producttextureController.Create)
	producttextureGroup.PATCH("/:id", producttextureController.Update)
	producttextureGroup.DELETE("/:id", producttextureController.Delete)

	//========== Smell Extract ==========
	smellextractController := controller.SmellExtracts{}
	smellextractGroup := r.Group("/smellextract")
	smellextractGroup.GET("", smellextractController.FindAll)
	smellextractGroup.GET("/:id", smellextractController.FindOne)
	smellextractGroup.POST("", smellextractController.Create)
	smellextractGroup.PATCH("/:id", smellextractController.Update)
	smellextractGroup.DELETE("/:id", smellextractController.Delete)

	//========== Smell Type ==========
	smelltypeController := controller.SmellType{}
	smelltypeGroup := r.Group("/smelltype")
	smelltypeGroup.GET("", smelltypeController.FindAll)
	smelltypeGroup.GET("/:id", smelltypeController.FindOne)
	smelltypeGroup.POST("", smelltypeController.Create)
	smelltypeGroup.PATCH("/:id", smelltypeController.Update)
	smelltypeGroup.DELETE("/:id", smelltypeController.Delete)

	//========== Smell Group ==========
	smellgroupController := controller.SmellGroup{}
	smellgroupGroup := r.Group("/smellgroup")
	smellgroupGroup.GET("", smellgroupController.FindAll)
	smellgroupGroup.GET("/:id", smellgroupController.FindOne)
	smellgroupGroup.POST("", smellgroupController.Create)
	smellgroupGroup.PATCH("/:id", smellgroupController.Update)
	smellgroupGroup.DELETE("/:id", smellgroupController.Delete)

	//========== Smell Group ==========
	smellController := controller.Smell{}
	smellGroup := r.Group("/smell")
	smellGroup.GET("", smellController.FindAll)
	smellGroup.GET("/:id", smellController.FindOne)
	smellGroup.POST("", smellController.Create)
	smellGroup.PATCH("/:id", smellController.Update)
	smellGroup.DELETE("/:id", smellController.Delete)

	//========== Units ==========
	unitsController := controller.Units{}
	unitsGroup := r.Group("/units")
	unitsGroup.GET("", unitsController.FindAll)
	unitsGroup.GET("/:id", unitsController.FindOne)
	unitsGroup.POST("", unitsController.Create)
	unitsGroup.PATCH("/:id", unitsController.Update)
	unitsGroup.DELETE("/:id", unitsController.Delete)

	//========== Test Size ==========
	testsizeController := controller.TestSize{}
	testsizeGroup := r.Group("/testsize")
	testsizeGroup.GET("", testsizeController.FindAll)
	testsizeGroup.GET("/:id", testsizeController.FindOne)
	testsizeGroup.POST("", testsizeController.Create)
	testsizeGroup.PATCH("/:id", testsizeController.Update)
	testsizeGroup.DELETE("/:id", testsizeController.Delete)

	//========== Product Formular ==========
	productformularController := controller.ProductFormula{}
	productformularGroup := r.Group("/productformular")
	productformularGroup.GET("", productformularController.FindAll)
	productformularGroup.GET("/:id", productformularController.FindOne)
	productformularGroup.POST("", productformularController.Create)
	productformularGroup.PATCH("/:id", productformularController.Update)
	productformularGroup.DELETE("/:id", productformularController.Delete)

	//========== Sales Team ==========
	salesteamController := controller.SalesTeam{}
	salesteamGroup := r.Group("/salesteam")
	salesteamGroup.GET("", salesteamController.FindAll)
	salesteamGroup.GET("/:id", salesteamController.FindOne)
	salesteamGroup.POST("", salesteamController.Create)
	salesteamGroup.PATCH("/:id", salesteamController.Update)
	salesteamGroup.DELETE("/:id", salesteamController.Delete)

	//========== Sales ==========
	salesController := controller.Sales{}
	salesGroup := r.Group("/sales")
	salesGroup.GET("", salesController.FindAll)
	salesGroup.GET("/:id", salesController.FindOne)
	salesGroup.POST("", salesController.Create)
	salesGroup.PATCH("/:id", salesController.Update)
	salesGroup.DELETE("/:id", salesController.Delete)

	//========== System Users ==========
	systemusersController := controller.SystemUsers{}
	systemuserGroup := r.Group("/systemuser")
	systemuserGroup.GET("", systemusersController.FindAll)
	systemuserGroup.GET("/:id", systemusersController.FindOne)
	systemuserGroup.POST("", systemusersController.Create)
	systemuserGroup.PATCH("/:id", systemusersController.Update)
	systemuserGroup.DELETE("/:id", systemusersController.Delete)

	//========== Users Group ==========
	usergroupController := controller.UsersGroup{}
	usergroupGroup := r.Group("/usergroup")
	usergroupGroup.GET("", usergroupController.FindAll)
	usergroupGroup.GET("/:id", usergroupController.FindOne)
	usergroupGroup.POST("", usergroupController.Create)
	usergroupGroup.PATCH("/:id", usergroupController.Update)
	usergroupGroup.DELETE("/:id", usergroupController.Delete)

	//========== Quotation ==========
	quotationController := controller.Quotation{}
	quotationGroup := r.Group("/quotation")
	quotationGroup.GET("", quotationController.FindAll)
	quotationGroup.GET("/:id", quotationController.FindOne)
	quotationGroup.POST("", quotationController.Create)
	quotationGroup.PATCH("/:id", quotationController.Update)
	quotationGroup.DELETE("/:id", quotationController.Delete)

	//========== Quotation Line ==========
	quotationlineController := controller.QuotationLine{}
	quotationlineGroup := r.Group("/quotationline")
	quotationlineGroup.GET("/:quotationId", quotationlineController.FindLineAll)
	quotationlineGroup.GET("/:quotationId/:lineId", quotationlineController.FindLineOne)
	quotationlineGroup.POST("", quotationlineController.Create)
	quotationlineGroup.PATCH("/:id", quotationlineController.Update)
	quotationlineGroup.DELETE("/:id", quotationlineController.Delete)

	//========== Quotation Line Option ==========
	quotationlineoptionController := controller.QuotationLineOption{}
	quotationlineoptionGroup := r.Group("/quotationlineoption")
	quotationlineoptionGroup.GET("/:quotationId/:lineId", quotationlineoptionController.FindLineOptionAll)
	quotationlineoptionGroup.GET("/:quotationId/:lineId/:optionId", quotationlineoptionController.FindLineOptionOne)
	quotationlineoptionGroup.POST("", quotationlineoptionController.Create)
	quotationlineoptionGroup.PATCH("/:id", quotationlineoptionController.Update)
	quotationlineoptionGroup.DELETE("/:id", quotationlineoptionController.Delete)
}
