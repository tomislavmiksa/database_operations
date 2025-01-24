package main

/*
go mod init github.com/tomislavmiksa/database_operations
go get gorm.io/gorm
go get gorm.io/driver/postgres
go mod tidy
*/

import (
	"fmt"
	"log"

	postgresdb "database_operations/postgressdb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "simple_bank"
)

func main() {
	// create the connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	fmt.Println("dsn:", dsn)

	// connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
		panic(err.Error())
	}
	fmt.Println("Database connected")

	// print all the tables in the connected database
	postgresdb.GetTableNames(db)

	// Migrate the schema
	// - creates the table if it does not exists
	// - modifies the schema as well
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	var pr postgresdb.Product
	var prs []postgresdb.Product

	// Create a single DB entry in the Table
	pr = postgresdb.InsertProduct(db, Product{Code: "D41", Price: 100})
	postgresdb.PrintProduct(pr)
	pr = postgresdb.InsertProduct(db, postgresdb.Product{Code: "D42", Price: 120})
	postgresdb.PrintProduct(pr)
	pr = postgresdb.InsertProduct(db, postgresdb.Product{Code: "D43", Price: 200})
	postgresdb.PrintProduct(pr)

	// Query for product with code D42, returns 1st result
	pr = postgresdb.GetProductByCode(db, "D43")
	postgresdb.PrintProduct(pr)

	// Modify a product or create if it does not exist
	pr = postgresdb.ModifyProduct(db, postgresdb.Product{Code: "D43", Price: 151})
	postgresdb.PrintProduct(pr)

	// Modify a product or create if it does not exist
	pr = postgresdb.ModifyProduct(db, postgresdb.Product{Code: "D44", Price: 220})
	postgresdb.PrintProduct(pr)

	// Query for product with code D42, returns 1st result
	prs = postgresdb.GetProductMoreExpensiveThen(db, 101)
	for _, pr := range prs {
		postgresdb.PrintProduct(pr)
	}

	// Query for all products; then delete all the products
	prs = postgresdb.GetAllProducts(db)
	for _, pr := range prs {
		postgresdb.PrintProduct(pr)
		db.Delete(&pr)
	}

}
