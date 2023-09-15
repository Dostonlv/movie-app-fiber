package router

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"movie/database"
	"movie/models"
	"movie/pkg/utils"
)

func User(c *fiber.Ctx) error {
	c.Accepts("application/json")

	collection := database.InitDB().Db.Collection("user")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user.Name = utils.TrimSpace(user.Name)
	user.Email = utils.TrimSpace(user.Email)
	user.Password = hashedPassword

	filter := bson.D{{Key: "email", Value: user.Email}}
	var result models.User
	err = collection.FindOne(c.Context(), filter).Decode(&result)
	if err == nil {
		return c.Status(401).SendString("email already in use")
	}

	insertionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	filter = bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	return c.Status(201).JSON(createdUser)

}
