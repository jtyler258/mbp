package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
 }

func NewDatabase(connection string) (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
	}, nil
}

func (d* Database) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := d.client.Connect(ctx)
	return err
}

func (d* Database) Disconnect(ctx context.Context) {
	d.client.Disconnect(ctx)
}