package main

import (
	"comments/configs"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Comment struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostId string             `json:"postId,omitempty" bson:"postId,omitempty"`
	Text   string             `json:"text,omitempty" bson:"text,omitempty"`
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic("Error loading config file")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	db := client.Database("go_search")

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		collection := db.Collection("comments")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		var comments []Comment
		cursor, err := collection.Find(ctx, Comment{PostId: c.Params("id")})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if err = cursor.All(ctx, &comments); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(comments)
	})
	app.Post("/api/comments", func(c *fiber.Ctx) error {
		collection := db.Collection("comments")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		comment := new(Comment)
		if err := c.BodyParser(comment); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		comment.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(ctx, comment)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(fiber.Map{"comment": comment})
	})

	app.Listen(":" + config.ServerPort)
}
