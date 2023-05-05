package database

import (
	"fmt"
	"github.com/yi-lin-1234/purchase-review-backend/config"
	"github.com/yi-lin-1234/purchase-review-backend/internal/model"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	//converts p from a string to a uint (using strconv.ParseUint)
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("parseUint failed")
	}

	// It constructs a database connection string (dsn), including host, port, user, password, database name, and sslmode. All these values are fetched from the configuration. The sslmode=disable part means the connection doesn't use SSL.
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connection opened to database 🟢")

	// Migrate the database
	DB.AutoMigrate(&model.Purchase{})
	fmt.Println("database migrated 🟢")
}
