package gorm_client

import (
	"fmt"

	"github.com/rysmaadit/go-template/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(app *app.Application) (db *gorm.DB, err error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.Config.DBUser,
		app.Config.DBPassword,
		app.Config.DBAddress,
		app.Config.DBPort,
		app.Config.DBName)

	db, err = gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
