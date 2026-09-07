package transfer

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	FindAccount(context.Context, string) (*Account, error)
	SetBalance(context.Context, string, float64) error
	CreateTransfer(context.Context, *Transfer) error
}
type repository struct{ db *mongo.Database }

func NewRepository(db *mongo.Database) Repository { return &repository{db: db} }
func (r *repository) FindAccount(ctx context.Context, id string) (*Account, error) {
	var a Account
	err := r.db.Collection("accounts").FindOne(ctx, bson.M{"_id": id}).Decode(&a)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}
func (r *repository) SetBalance(ctx context.Context, id string, b float64) error {
	_, err := r.db.Collection("accounts").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"balance": b}})
	return err
}
func (r *repository) CreateTransfer(ctx context.Context, t *Transfer) error {
	_, err := r.db.Collection("transactions").InsertOne(ctx, t)
	return err
}
