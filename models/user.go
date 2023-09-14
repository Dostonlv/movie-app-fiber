package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
