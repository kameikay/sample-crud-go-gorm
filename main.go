package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// // CREATE
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// CREATE BATCH
	// products := []Product{
	// 	{Name: "Notebook", Price: 10000.00},
	// 	{Name: "Mouse", Price: 400.00},
	// 	{Name: "Keyboard", Price: 600.00},
	// }
	// db.Create(&products)

	// SELECT ONE
	// var product Product

	// db.First(&product, 2)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// SELECT ALL
	// var products []Product
	// db.Find(&products)

	// LIMIT AND PAGINATION
	// db.Limit(2).Offset(1).Find(&products)

	// WHERE
	// db.Where("NAME LIKE ?", "%k%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// UPDATE
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Name"
	// db.Save(&p)

	// db.First(&p, 1)
	// fmt.Println(p)

	// DELETE
	var p2 Product
	db.First(&p2, 1)
	db.Delete(&p2)
}
