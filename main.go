package main

import (
	"gogorm/controller"
)

// "fmt"
// "gogorm/service"

// "github.com/spf13/viper"
// "gorm.io/driver/mysql"
// "gorm.io/gorm"

func main() {

	controller.StartServer()

	// service := service.NewShowDataService(db)
	// service.PrintAllCountries()

	// countries := []model.Conuntry{}
	// result := db.Find(&countries)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// for _, item := range countries {
	// 	fmt.Printf("%v\n", item)
	// }

	// landmarks := []model.Landmark{}
	// result := db.Joins("Country").Find(&landmarks)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// for _, item := range landmarks {
	// 	fmt.Printf("%v\n", item)
	// }

	// crate
	// country := model.Country{Name: "ประเทศปากช่อง"}
	// result := db.Create(&country)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// if result.RowsAffected > 0 {
	// 	fmt.Println("Insert Completed")
	// }

	//delete
	//ไปดูที่โมเดลว่ามันผูกกับ table ไหน แล้วลบจาก table นั้นโดยส่ง id เข้าไป
	//country := model.Country{Name: "ประเทศปากช่อง"}
	// result := db.Delete(&model.Country{}, 122)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// if result.RowsAffected > 0 {
	// 	fmt.Println("delete Completed")
	// }

	//delete by condition
	//result := db.Delete(&model.Country{}, "Name = ?", "ปากช่อง")
	// result := db.Where("Name = ?", "ปากช่อง").Delete(&model.Country{})
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// if result.RowsAffected > 0 {
	// 	fmt.Println("delete Completed")
	// }

	//update
	//1.qeury frist
	//country := model.Country{}
	//result := db.Find(&country, 19)
	//update 1 column
	//result := db.Model(&country).Where("Name = ?", "ประเทศน้องจุ๊ก").Update("Name", "ประเทศน้องจุ๊กvbvbvvb")
	// result := db.Model(&model.Landmark{}).Where("Name = ?", "ประเทศ").Updates(model.Landmark{Name: "น้ำ", Detail: "งง"})
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// //2.update
	// // country.Name = "ประเทศน้องจุ๊กvbvbvvb"
	// result = db.Save(&country)
	// if result.Error != nil {
	// 	panic(result.Error)
	// } else {
	// 	fmt.Println("update Completed")
	// }

	//find
	//qeury no condition
	// landmarks := []model.Landmark{}
	// result := db.Find(&landmarks)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// for _, v := range landmarks {
	// 	fmt.Println(v)
	// 	fmt.Println("===========================")
	// }

	//where
	//qeury with condition
	// landmarks := []model.Landmark{}
	// result := db.Where("Country = ? and Name like ?", 1, "%ปาก%").Find(&landmarks)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// for _, v := range landmarks {
	// 	fmt.Println(v)
	// 	fmt.Println("===========================")
	// }

	//take
	// landmark := model.Landmark{}
	// result := db.Take(&landmark, 4)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(landmark)
	// fmt.Println("===========================")

	//first
	//จะค้นหาตัวแรกที่เจอในดาต้าเบส
	// landmark := model.Landmark{}
	// result := db.Where("Country = ?", 3).First(&landmark)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(landmark)
	// fmt.Println("===========================")

	//last
	// landmark := model.Landmark{}
	// result := db.Where("Country = ?", 1).Last(&landmark)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(landmark)
	// fmt.Println("===========================")

	//order
	// landmarks := []model.Landmark{}
	// result := db.Where("Country = ?", 1).Order("name desc").Limit(3).Find(&landmarks)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// for _, v := range landmarks {
	// 	fmt.Println(v)
	// 	fmt.Println("===========================")
	// }

	// type dist struct {
	// 	Id uint `gorm:"column:idx"`
	// }
	// results := []dist{}
	// result := db.Model(model.Country{}).Distinct("idx").Find(&results)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(results)

}

// func main() {
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%v\n", viper.GetString("mysql.dsn"))
// 	dsn := viper.GetString("mysql.dsn")
// 	dialactor := mysql.Open(dsn)

// 	db, err := gorm.Open(dialactor)
// 	if err != nil {
// 		panic(err)
// 	}
// 	println("connection success")

// 	// countries := []model.Country{}
// 	// result := db.Find(&countries)
// 	// if result.Error != nil{
// 	// 	panic(result.Error)
// 	// }
// 	// for _, v := range countries{
// 	// 	fmt.Printf("%v\n",v)
// 	// }

// 	landmarks := []model.Landmark{}
// 	result := db.Joins("Country").Find(&landmarks)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	for _, v := range landmarks {
// 		fmt.Printf("%v\n", v)
// 	}

// }
