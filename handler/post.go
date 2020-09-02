package handler

import (
	"context"

	"github.com/FranciscoMendes10866/api-go/prisma/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func CreatePost(c *fiber.Ctx) {
	ctx := context.Background()
	prisma := db.NewClient()
	err := prisma.Connect()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := prisma.Disconnect()
		if err != nil {
			panic(err)
		}
	}()

	type request struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Published bool   `json:"published"`
	}

	body := new(request)
	c.BodyParser(body)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authEmail := claims["email"].(string)

	create, err := prisma.Post.CreateOne(
		db.Post.Title.Set(body.Title),
		db.Post.Content.Set(body.Content),
		db.Post.Published.Set(body.Published),
		db.Post.Author.Link(
			db.User.Email.Equals(authEmail),
		),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(create)
}
