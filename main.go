package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "belajar_golang"
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var db *gorm.DB
var err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"sslmode=disable dbname=%s password=%s",
		host, port, user, dbname, password)
	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	p1 := Person{FirstName: "John", LastName: "Doe"}
	p2 := Person{FirstName: "Jane", LastName: "Smith"}

	var p3 Person
	db.First(&p3)

	fmt.Println(p1.FirstName)
	fmt.Println(p2.LastName)
	fmt.Println(p3.LastName)

	r := gin.Default()
	r.GET("/", HelloWorld)
	r.GET("/people", GetPeople)
	r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)

	r.Run(":8081")
}

func HelloWorld(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func GetPeople(ctx *gin.Context) {
	var people []Person

	err = db.Find(&people).Error
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	ctx.JSON(200, people)
}

func GetPerson(ctx *gin.Context) {
	var person Person
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
	var person Person
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
	var person Person
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
	id := ctx.Params.ByName("id")
	var person Person
	err = db.Where("id = ?", id).Delete(&person).Error
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	ctx.JSON(200, gin.H{"id #" + id: "deleted"})
}
