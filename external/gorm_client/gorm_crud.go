package gorm_client

import (
	"github.com/rysmaadit/go-template/app"
	"gorm.io/gorm"
)

func InitMigration() (db *gorm.DB, err error) {
	db, err = Connect(app.Init())
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(Movie{})
	return db, nil
}

func CreateMovie(movie Movie) (err error) {
	db, err := InitMigration()
	if err != nil {
		return err
	}
	db.Create(&movie)
	return nil
}

func ReadAll() (movie []Movie, err error) {
	db, _ := InitMigration()

	db.Find(&movie)
	return movie, nil
}

func ReadMovie(slug string) (movie Movie, err error) {
	db, err := InitMigration()
	if err != nil {
		return Movie{}, err
	}
	movie = Movie{Slug: slug}
	res := db.First(&movie, "slug=?", slug)
	if res.Error != nil {
		return Movie{}, err
	}
	return movie, nil
}

func UpdateMovie(slug string, movie Movie) (err error) {
	db, err := InitMigration()
	if err != nil {
		return err
	}
	db.Model(&movie).Where("slug=?", slug).Updates(&movie)
	return nil
}

func DeleteMovie(slug string) (err error) {
	db, err := InitMigration()
	if err != nil {
		return err
	}
	var movie Movie
	db.Model(&movie).Where("slug=?", slug).Delete(&movie)
	return nil
}
