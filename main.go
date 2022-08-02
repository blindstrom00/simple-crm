package main

import (
	"fmt"

	"github.com/blindstrom00/simple-crm/database"
	"github.com/blindstrom00/simple-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func routeSetup(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)    // mutliple leads
	app.Get("/api/v1/lead/:id", lead.GetLead) // single lead
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection to database is now open")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database has been migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	routeSetup(app)
	app.Listen(8080)
	defer database.DBconn.Close()
}
