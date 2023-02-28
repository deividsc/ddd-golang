package purchase

import (
	coffeeco "coffeeco-monolithic/internal"
	"coffeeco-monolithic/internal/payment"
	"coffeeco-monolithic/internal/store"
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

func NewMongoRepository(ctx context.Context, connectionStrinng string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStrinng))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (r *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	_, err := r.purchases.InsertOne(ctx, purchase)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}

	return nil
}

type mongoPurchase struct {
	id                uuid.UUID
	store             store.Store
	productToPurchase []coffeeco.Product
	total             money.Money
	paymentMeans      payment.Means
	timeOfPurchase    time.Time
	cardToken         *string
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		id:                p.id,
		store:             p.Store,
		productToPurchase: p.ProductToPurchase,
		total:             p.total,
		paymentMeans:      p.PaymentMeans,
		timeOfPurchase:    p.timeOfPurchase,
		cardToken:         p.cardToken,
	}
}
