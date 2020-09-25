package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go_gin_api/router"
)

func main() {
	r := gin.Default()
	router.Route(r)
	r.Run(":8081")
}
