package postgressdb

/*
go mod init database_operations/postgressdb
go get gorm.io/gorm
go get gorm.io/driver/postgres
go mod tidy
*/

import (
	"fmt"

	"gorm.io/gorm"
)

// table definition in gorm
// gorm.Model definition
//
//	type Model struct {
//		ID        uint           `gorm:"primaryKey"`
//		CreatedAt time.Time
//		UpdatedAt time.Time
//		DeletedAt gorm.DeletedAt `gorm:"index"`
//	}
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// get the table names in the database
// this is helper function, to not be used
func GetTableNames(db *gorm.DB) {
	var tables []string
	if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	fmt.Println(" List of Tables: ")
	fmt.Println("-------------------------------------------------")
	fmt.Println(tables)
}

// print the product struct
func PrintProduct(p Product) {
	fmt.Println("ID      : ", p.ID)
	fmt.Println("Code    : ", p.Code)
	fmt.Println("Price   : ", p.Price)
	fmt.Println("Created : ", p.CreatedAt)
	fmt.Println("Updated : ", p.UpdatedAt)
	fmt.Println("Deleted : ", p.DeletedAt)
	fmt.Println("-------------------------------------------------")
}

func InsertProduct(db *gorm.DB, prod Product) Product {
	// Create a single DB entry in the Table
	fmt.Println("")
	fmt.Println("Creating a single row in the table by the Product Code")
	result := db.Create(&prod)
	if result.Error != nil {
		fmt.Println(result.Error)
		return prod
	} else {
		fmt.Println("Successfully created row", prod)
		fmt.Println("Rows Affected    : ", result.RowsAffected)
		fmt.Println("Read Error       : ", result.Error)
		fmt.Println("-------------------------------------------------")
		return prod
	}
}

func ModifyProduct(db *gorm.DB, prod Product) Product {
	// Modify a DB entry in the Table
	fmt.Println("")
	fmt.Println("Modifying the product in the table!")
	result := db.Save(&prod)
	if result.Error != nil {
		fmt.Println(result.Error)
		return prod
	} else {
		fmt.Println("Successfully created row", prod)
		fmt.Println("Rows Affected    : ", result.RowsAffected)
		fmt.Println("Read Error       : ", result.Error)
		fmt.Println("-------------------------------------------------")
		return prod
	}
}

func GetAllProducts(db *gorm.DB) []Product {
	// Read a single DB entry in the Table with Code as Filter
	fmt.Println("")
	fmt.Println("Getting a all rows in the table!")
	var products []Product
	result := db.Find(&products) // product --> type of *gorm.DB
	if result.Error != nil {
		fmt.Println(result.Error)
		return products
	} else {
		fmt.Println("Rows Affected    : ", result.RowsAffected)
		fmt.Println("Read Error       : ", result.Error)
		fmt.Println("-------------------------------------------------")
		return products
	}
}

func GetProductByCode(db *gorm.DB, code string) Product {
	// Read a single DB entry in the Table with Code as Filter
	fmt.Println("")
	fmt.Println("Getting a single row in the table with code: ", code)
	product := Product{}
	result := db.First(&product, "code = ?", code) // product --> type of *gorm.DB
	if result.Error != nil {
		fmt.Println(result.Error)
		return product
	} else {
		fmt.Println("Rows Affected    : ", result.RowsAffected)
		fmt.Println("Read Error       : ", result.Error)
		fmt.Println("-------------------------------------------------")
		return product
	}
}

func GetProductMoreExpensiveThen(db *gorm.DB, price uint) []Product {
	// Read a DB entries in the Table with Code as Filter
	fmt.Println("")
	fmt.Println("Getting a rows in the table where price higher then: ", price)
	var products []Product
	result := db.Where("price > ?", price).Find(&products) // product --> type of *gorm.DB
	if result.Error != nil {
		fmt.Println(result.Error)
		return products
	} else {
		fmt.Println("Rows Affected    : ", result.RowsAffected)
		fmt.Println("Read Error       : ", result.Error)
		fmt.Println("-------------------------------------------------")
		return products
	}
}
