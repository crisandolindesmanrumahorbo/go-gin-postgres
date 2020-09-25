package util

import (
	"github.com/jinzhu/gorm"
	"go_gin_api/model"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Person{})
}
