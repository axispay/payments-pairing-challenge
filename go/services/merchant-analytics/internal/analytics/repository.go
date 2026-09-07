package analytics

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type Repository interface {
	CreateTransaction(context.Context, *Transaction) error
	FindTransactions(context.Context, string, time.Time, time.Time) ([]bson.M, error)
}
type repository struct{ collection *mongo.Collection }

func NewRepository(db *mongo.Database) Repository { return &repository{db.Collection("transactions")} }

func (r *repository) CreateTransaction(ctx context.Context, transaction *Transaction) error {
	_, err := r.collection.InsertOne(ctx, transaction)
	return err
}

func (r *repository) FindTransactions(ctx context.Context, merchantID string, start, end time.Time) ([]bson.M, error) {
	cur, err := r.collection.Find(ctx, bson.M{"merchantId": merchantID, "createdAt": bson.M{"$gte": start, "$lt": end}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []bson.M
	for cur.Next(ctx) {
		var tx bson.M
		if cur.Decode(&tx) == nil {
			out = append(out, tx)
		}
	}
	return out, cur.Err()
}
