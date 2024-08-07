package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/dto"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

// FindOne gets a document according to CheckPermissionDto model.
func (m *MongoClient) FindOne(c dto.CheckPermissionDto) (model.Key, error) {
	collection := m.client.Database("key_manager").Collection("keys")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result model.Key
	err := collection.FindOne(ctx, bson.M{"apiKey": c.ApiKey}).Decode(&result)
	return result, err
}

// InsertOne inserts an item on key_manager database.
func (m *MongoClient) InsertOne(i interface{}) (interface{}, error) {
	collection := m.client.Database("key_manager").Collection("keys")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, i)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
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
