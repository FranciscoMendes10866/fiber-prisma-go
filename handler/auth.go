package handler

import (
	"context"

	"github.com/FranciscoMendes10866/api-go/prisma/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func GetUsers(c *fiber.Ctx) {
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

	users, err := prisma.User.FindMany().Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(users)
}

func CreateUser(c *fiber.Ctx) {
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
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	body := new(request)
	c.BodyParser(body)

	create, err := prisma.User.CreateOne(
		db.User.Email.Set(body.Email),
		db.User.Name.Set(body.Name),
		db.User.Age.Set(body.Age),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(create)
}

func LoginUser(c *fiber.Ctx) {
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
		Email string `json:"email"`
	}

	body := new(request)
	c.BodyParser(body)

	login, err := prisma.User.FindOne(
		db.User.Email.Equals(body.Email),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = login.Email
	// Generate encoded token and send it as response.
	loginToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
	// Response
	c.JSON(fiber.Map{"token": loginToken})
}
