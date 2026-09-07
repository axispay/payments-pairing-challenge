package account

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	Create(context.Context, *Account) error
	FindByID(context.Context, string) (*Account, error)
	SetBalance(context.Context, string, float64) error
}
type repository struct{ collection *mongo.Collection }

func NewRepository(db *mongo.Database) Repository {
	return &repository{collection: db.Collection("accounts")}
}
func (r *repository) Create(ctx context.Context, account *Account) error {
	_, err := r.collection.InsertOne(ctx, account)
	return err
}

func (r *repository) FindByID(ctx context.Context, id string) (*Account, error) {
	var a Account
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&a)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}
func (r *repository) SetBalance(ctx context.Context, id string, balance float64) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"balance": balance}})
	return err
}
