package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	Id     uint   `json:"id"`
	PostId uint   `json:"postid"`
	Text   string `json:"text"`
}

func main() {
	dsn := "host=localhost user=postgres password=eljc102030 dbname=comments_ms port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Comment{})

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		var comments []Comment

		db.Find(&comments, "post_id = ?", c.Params("id"))

		return c.JSON(comments)
	})

	app.Get("/api/comments", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/comment", func(c *fiber.Ctx) error {
		var comment Comment

		if err := c.BodyParser(&comment); err != nil {
			return err
		}

		db.Create(&comment)

		return c.JSON(comment)
	})

	app.Listen(":3001")
}
