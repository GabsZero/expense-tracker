package repositories

import (
	"fmt"
	"os"

	"github.com/gabszero/expense-tracker/pkg/infra/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db gorm.DB
}

func (databaseStruct *Database) startDatabase() *Database {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)
	connectedDatabase, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("ops")
	}

	database := Database{
		db: *connectedDatabase,
	}

	return &database
}

func (d *Database) AutoMigrate() {
	db := d.startDatabase()
	db.db.AutoMigrate(&models.Expense{})
}
