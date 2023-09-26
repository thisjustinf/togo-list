package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
const (
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
)

func InitDB(){
	var err error
    db, err = gorm.Open("postgres", "")

    if err != nil {
        panic("failed to connect database")
    }
}