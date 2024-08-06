package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

// Disconnect closes the socket connection.
func (m *MongoClient) Disconnect() error {
	if err := m.client.Disconnect(m.ctx); err != nil {
		return err
	}
	return nil
}

// Ping send to server a ping command to check the connection.
func (m *MongoClient) Ping() error {
	err := m.client.Ping(m.ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

// NewMongoClient create an instance of MongoClient struct.
func NewMongoClient(ctx context.Context) (*MongoClient, error) {
	connUri := func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return os.Getenv("CONN_STRING_MONGODB_DEBUG")
		}
		return os.Getenv("CONN_STRING_MONGODB_RELEASE")
	}()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connUri))
	if err != nil {
		return nil, err
	}
	return &MongoClient{client: client, ctx: ctx}, nil
}
