package reports

import (
	"fmt"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var EthReports *types.EthReports

type DatabaseManager struct {
	DB *gorm.DB
}

func Initialize(dbPass, dbURL, dbUser string) (*DatabaseManager, error) {
	dbm := DatabaseManager{}
	db, err := OpenDBConnection(dbPass, dbURL, dbUser)
	if err != nil {
		return nil, err
	}
	dbm.DB = db
	dbm.RunMigrations()
	return &dbm, nil
}

func (dbm *DatabaseManager) RunMigrations() {
	dbm.DB.AutoMigrate(EthReports)
}

// OpenDBConnection is used to create a database connection
func OpenDBConnection(dbPass, dbURL, dbUser string) (*gorm.DB, error) {
	if dbUser == "" {
		dbUser = "postgres"
	}
	// look into whether or not we wil disable sslmode
	dbConnURL := fmt.Sprintf("host=%s port=5432 user=%s dbname=reports password=%s", dbURL, dbUser, dbPass)
	db, err := gorm.Open("postgres", dbConnURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func OpenTestDBConnection(dbPass string) (*gorm.DB, error) {
	dbConnURL := fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres dbname=reports password=%s", dbPass)
	db, err := gorm.Open("postgres", dbConnURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDBConnection is used to close a db
func CloseDBConnection(db *gorm.DB) {
	db.Close()
}
