package auction

import (
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionMongo struct {
	Id          string                   `bson:"_id"`
	ProductName string                   `bson:"product_name"`
	Category    string                   `bson:"category"`
	Description string                   `bson:"description"`
	Condition   auction.ProductCondition `bson:"condition"`
	Status      auction.AuctionStatus    `bson:"status"`
	Timestamp   int64                    `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}
