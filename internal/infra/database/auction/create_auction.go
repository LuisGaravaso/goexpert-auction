package auction

import (
	"context"
	"os"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *AuctionRepository) CreateAuction(ctx context.Context, auc auction.Auction) (*auction.Auction, *internal_errors.InternalError) {
	auctionMongo := AuctionMongo{
		Id:          auc.Id,
		ProductName: auc.ProductName,
		Category:    auc.Category,
		Description: auc.Description,
		Condition:   auc.Condition,
		Status:      auc.Status,
		Timestamp:   auc.Timestamp.Unix(),
	}

	if _, err := r.Collection.InsertOne(ctx, auctionMongo); err != nil {
		logger.Error("error creating auction", err)
		return nil, internal_errors.NewInternalServerError("error creating auction")
	}

	go func() {
		select {
		case <-time.After(getAuctionInterval()):
			update := bson.M{
				"$set": bson.M{"status": auction.Completed},
			}
			filter := bson.M{"_id": auc.Id}

			_, err := r.Collection.UpdateOne(ctx, filter, update)
			if err != nil {
				logger.Error("error trying to complete the auction", err)
				return
			}
		}
	}()

	return &auc, nil
}

func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return 5 * time.Minute
	}

	return duration
}
