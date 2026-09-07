package payment

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	Create(context.Context, *Payment) error
}
type repository struct{ collection *mongo.Collection }

func NewRepository(db *mongo.Database) Repository { return &repository{db.Collection("payments")} }
func (r *repository) Create(ctx context.Context, p *Payment) error {
	res, err := r.collection.InsertOne(ctx, p)
	if err == nil {
		p.ID = res.InsertedID
	}
	return err
}
