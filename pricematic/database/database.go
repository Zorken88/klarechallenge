package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/KlareTeam/interview-challenges/go/pricematic/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbport, _ := strconv.Atoi(os.Getenv("DBPORT"))

	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DBHOST"), dbport, os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.Product{}, &models.Price{})
	dataSeed(db)

	Database = DbInstance{
		Db: db,
	}
}

func dataSeed(db *gorm.DB) {

	// Open json file data.json
	jsonFile, err := os.Open(os.Getenv("JSON_FILE"))

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	products := []models.Product{}
	price := models.Price{}

	json.Unmarshal([]byte(byteValue), &products)

	for _, product := range products {

		price.Value = product.ActualPrice

		product.Prices = append(product.Prices, price)
		db.Create(&product)
	}

}
