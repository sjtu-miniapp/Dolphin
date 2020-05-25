package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

// REF: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
func MongoConn(host, user, pass string, port int) (*mongo.Client, error) {
	// Set client options
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	clientOptions := options.Client().ApplyURI(uri)
	clientOptions.Auth = &options.Credential{
		Username: user,
		Password: pass,
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	return client, err
}
