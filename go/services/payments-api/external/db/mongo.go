package db

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func Connect(uri string) (*Mongo, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	name := "payments"
	if i := strings.LastIndex(uri, "/"); i >= 0 && i+1 < len(uri) {
		name = strings.Split(uri[i+1:], "?")[0]
	}
	return &Mongo{client: client, db: client.Database(name)}, nil
}

func (m *Mongo) DB() *mongo.Database   { return m.db }
func (m *Mongo) Client() *mongo.Client { return m.client }
func (m *Mongo) Disconnect()           { _ = m.client.Disconnect(context.Background()) }
