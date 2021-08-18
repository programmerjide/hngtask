package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type contact struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Website string `json:"website" form:"website"`
	Subject string `json:"subject" form:"subject"`
	Message string `json:"message" form:"message"`
}

func home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":   "Welcome to my Resume",
		"Message": "Olajide",
	})
}

func processData(ctx *fiber.Ctx) error {
	con := new(contact)
	err := ctx.BodyParser(con)
	if err != nil {
		log.Fatal(err)
	}

	return ctx.Render("index", fiber.Map{
		"senderName":  con.Name,
		"senderEmail": con.Email,
	})
}

func handleFunc(app *fiber.App) {
	app.Post("/", processData)
	app.Get("/", home)
}

func main() {

	//port := os.Getenv("PORT")
	port := os.Getenv("PORT")
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	handleFunc(app)

	log.Fatal(app.Listen(":" + port))
}
