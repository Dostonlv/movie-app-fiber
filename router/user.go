package router

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"movie/database"
	"movie/models"
)

func User(c *fiber.Ctx) error {
	c.Accepts("application/json")

	collection := database.InitDB().Db.Collection("user")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	insertionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Employee
	createdEmployee := &models.User{}
	createdRecord.Decode(createdEmployee)

	// return the created Employee in JSON format
	return c.Status(201).JSON(createdEmployee)

}
