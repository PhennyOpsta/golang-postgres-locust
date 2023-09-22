package main

import (
	"database/sql"
	"fmt"
	"go-api/controller/api"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func apiPath(c *fiber.Ctx) error {
	return c.SendString("Hello, World! API")
}

func setupRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/", helloWorld)
	app.Get("/api", apiPath)
	app.Post("/api/plus", api.GetPlus)
	app.Post("/api/minus", api.GetMinus)
	app.Post("/api/multiple", api.GetMultiple)
	app.Post("/api/divide", api.GetDivide)
	app.Post("/api/product/:id", func(c *fiber.Ctx) error {
		return api.GetProductDetail(c, db)
	})

}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("connect postgres error:", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("ping postgres error")
	}

	app := fiber.New()
	setupRoutes(app, db)

	app.Listen(":3000")
}
