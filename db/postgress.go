package db

import (
	"fmt"
	"log"
	"os"

	"github.com/AyoOluwa-Israel/invoice-api/config"
	"github.com/AyoOluwa-Israel/invoice-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func NewConnection(config *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUserName, config.DBUserPassword, config.DBName,
	)

	devDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DevDBHost, config.DevDBPort, config.DevDBUserName, config.DevDBUserPassword, config.DevDBName,
	)

	var url string
	if config.AppEnv == "development" {
		url = devDsn
	} else {
		url = dsn
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	fmt.Println("🚀 Connected Successfully to the Database")

	db.AutoMigrate(&models.User{}, &models.Invoice{}, &models.PaymentInformation{})

	Database = DbInstance{
		Db: db,
	}

}
