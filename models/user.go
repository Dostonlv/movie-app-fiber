package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
