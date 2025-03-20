package auction

import (
	"context"
	"errors"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction.Auction, *internal_errors.InternalError) {
	filter := bson.M{"_id": id}
	var auctionMongo AuctionMongo
	if err := r.Collection.FindOne(ctx, filter).Decode(&auctionMongo); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("auction not found with id: "+id, err)
			return &auction.Auction{}, internal_errors.NewNotFoundError("auction not found with id: " + id)
		}

		logger.Error("error finding auction by id: "+id, err)
		return &auction.Auction{}, internal_errors.NewInternalServerError("error finding auction by id: " + id)
	}

	return &auction.Auction{
		Id:          auctionMongo.Id,
		ProductName: auctionMongo.ProductName,
		Category:    auctionMongo.Category,
		Description: auctionMongo.Description,
		Condition:   auctionMongo.Condition,
		Status:      auctionMongo.Status,
		Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
	}, nil
}

func (r *AuctionRepository) FindAuctions(ctx context.Context, status auction.AuctionStatus, category string, productName string) ([]auction.Auction, *internal_errors.InternalError) {
	filter := bson.M{}
	if status != 0 {
		filter["status"] = status
	}
	if category != "" {
		filter["category"] = category
	}
	if productName != "" {
		filter["product_name"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("error finding auctions", err)
		return nil, internal_errors.NewInternalServerError("error finding auctions")
	}
	defer cursor.Close(ctx)

	var auctions []auction.Auction
	for cursor.Next(ctx) {
		var auctionMongo AuctionMongo
		if err := cursor.Decode(&auctionMongo); err != nil {
			logger.Error("error decoding auction", err)
			return nil, internal_errors.NewInternalServerError("error decoding auction")
		}

		auctions = append(auctions, auction.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctions, nil
}
