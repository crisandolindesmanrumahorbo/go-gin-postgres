package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"go_gin_api/model"
	"go_gin_api/util"
	"net/http"
)

var err error

func HelloWorld(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func GetPeople(ctx *gin.Context) {
	var people []model.Person
	db := util.DBConnect()
	defer db.Close()

	err = db.Find(&people).Error
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	ctx.JSON(200, people)
}

func GetPerson(ctx *gin.Context) {
	db := util.DBConnect()
	defer db.Close()
	var person model.Person
	id := ctx.Params.ByName("id")
	err = db.Where("id = ?", id).First(&person).Error
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	ctx.JSON(200, person)
}

func CreatePerson(ctx *gin.Context) {
	var person model.Person
	db := util.DBConnect()
	defer db.Close()
	err = ctx.BindJSON(&person)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		fmt.Println(err)
		return
	}
	db.Create(&person)
	ctx.JSON(http.StatusCreated, person)
}

func UpdatePerson(ctx *gin.Context) {
	var person model.Person
	db := util.DBConnect()
	defer db.Close()
	id := ctx.Params.ByName("id")

	err = db.Where("id = ?", id).First(&person).Error
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	err = ctx.BindJSON(&person)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		fmt.Println(err)
		return
	}
	db.Save(&person)
	ctx.JSON(200, person)
}

func DeletePerson(ctx *gin.Context) {
	db := util.DBConnect()
	defer db.Close()
	id := ctx.Params.ByName("id")
	var person model.Person
	err = db.Where("id = ?", id).Delete(&person).Error
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	ctx.JSON(200, gin.H{"id #" + id: "deleted"})
}
