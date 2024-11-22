package controller

import (
	"gogorm/model"
	"gogorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewCountriesController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	GetData := router.Group("/GET")
	{
		GetData.GET("/AllCountries", GetAllCountries)
		GetData.GET("/nameCountry", GetByNameCountry)
		GetData.GET("/:idx", GetByID)
	}
	DeleteData := router.Group("/Delete")
	{
		DeleteData.DELETE("/:idx", DeleteByID)
	}
	CreateData := router.Group("/Insert")
	{
		CreateData.POST("/", InsertData)

	}

}
func InsertData(ctx *gin.Context) {
	//presentation call service
	json := model.Country{}
	ctx.ShouldBindJSON(&json)
	service := service.NewShowDataService(db)
	result := service.InsertCountry(json)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, result.Error)

	}
	ctx.JSON(http.StatusOK, gin.H{
		"rowsaffected": result.RowsAffected,
		"name":         json.Name,
	})

}
func DeleteByID(ctx *gin.Context) {
	//presentation call service
	idxStr := ctx.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	service := service.NewShowDataService(db)
	result := service.DeleteByID(idx)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, result.Error)
	}
	ctx.JSON(http.StatusOK, result.RowsAffected)

}
func GetByID(ctx *gin.Context) {
	//presentation call service
	idxStr := ctx.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	service := service.NewShowDataService(db)
	result, err := service.GetCountryByID(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, result)

}
func GetByNameCountry(ctx *gin.Context) {
	//presentation call service
	nameCountry := ctx.Query("nameCountry")
	service := service.NewShowDataService(db)
	result, err := service.GetAllCountriesByName(nameCountry)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, result)

}
func GetAllCountries(ctx *gin.Context) {
	//presentation call service
	service := service.NewShowDataService(db)
	result, err := service.PrintAllCountries()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, result)

}
