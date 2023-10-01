package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"fmt"
	"github.com/adindazenn/assignment3-03/assignment3/model"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "root"
	dbname		= "postgres"
)

func InitDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
	db.Debug().AutoMigrate(model.Data{})
    
    return db, nil
}
