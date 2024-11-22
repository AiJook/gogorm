package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dsn := viper.GetString("mysql.dsn")
	//creat connection
	dialetor := mysql.Open(dsn)

	db, err := gorm.Open(dialetor)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Success")
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK,
			"API is working....")
	})

	NewCountriesController(router, db)
	router.Run()

}
