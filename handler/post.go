package handler

import (
	"context"

	"github.com/FranciscoMendes10866/api-go/prisma/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func FindUserPosts(c *fiber.Ctx) {
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

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authEmail := claims["email"].(string)

	posts, err := prisma.Post.FindMany(
		db.Post.Author.Where(
			db.User.Email.Equals(authEmail),
		),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(posts)
}

func FindSinglePost(c *fiber.Ctx) {
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

	PostID := c.Params("id")

	post, err := prisma.Post.FindOne(
		db.Post.ID.Equals(PostID),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(post)
}

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
		db.Post.Author.Link(
			db.User.Email.Equals(authEmail),
		),
		db.Post.Title.Set(body.Title),
		db.Post.Content.Set(body.Content),
		db.Post.Published.Set(body.Published),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(create)
}
