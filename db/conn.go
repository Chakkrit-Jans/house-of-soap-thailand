package db

import (
	"log"
	"os"

	"hst-api/model"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(
		//mysql.Open(dsn),
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&model.AppConfig{},
		&model.Title{},
		&model.UsersGroup{},
		&model.SystemAuthorized{},
		&model.SystemUsers{},
		&model.UsersGroupDetail{},
		&model.UsersGroupAuthorized{},
		&model.UserPermissions{},
		&model.UsersLogs{},
		&model.SalesTeam{},
		&model.Sales{},
		&model.Vat{},
		&model.CustomerGroup{},
		&model.CustomerType{},
		&model.Customers{},
		&model.Discount{},
		&model.UserApprovedDiscount{},
		&model.Units{},
		&model.TestSize{},
		&model.SmellExtracts{},
		&model.SmellGroup{},
		&model.SmellType{},
		&model.Smell{},
		&model.ProductTexture{},
		&model.Claim{},
		&model.Howto{},
		&model.Color1{},
		&model.Color2{},
		&model.Color3{},
		&model.Product{},
		&model.ProductType{},
		&model.ProductFormula{},
		&model.Cost{},
		&model.Quotation{},
		&model.QuotationLine{},
		&model.QuotationLineOption{},
		&model.NumberSequences{},
	)
}
