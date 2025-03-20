package bid

import (
	"context"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/bid"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid.Bid, *internal_errors.InternalError) {
	cursor, err := r.Collection.Find(ctx, bson.M{"auction_id": auctionId})

	errString := "error finding bid by auction id " + auctionId
	if err != nil {
		logger.Error(errString, err)
		return nil, internal_errors.NewInternalServerError(errString)
	}

	var bids []BidMongo
	if err := cursor.All(ctx, &bids); err != nil {
		logger.Error(errString, err)
		return nil, internal_errors.NewInternalServerError(errString)
	}

	var bidEntities []bid.Bid
	for _, bidMongo := range bids {
		bidEntities = append(bidEntities, bid.Bid{
			Id:        bidMongo.Id,
			UserId:    bidMongo.UserId,
			AuctionId: bidMongo.AuctionId,
			Amount:    bidMongo.Amount,
			Timestamp: time.Unix(bidMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (r *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid.Bid, *internal_errors.InternalError) {
	filter := bson.M{"auction_id": auctionId}
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bidMongo BidMongo
	if err := r.Collection.FindOne(ctx, filter, opts).Decode(&bidMongo); err != nil {
		errString := "error finding winning bid by auction id " + auctionId
		logger.Error(errString, err)
		return nil, internal_errors.NewInternalServerError(errString)
	}

	return &bid.Bid{
		Id:        bidMongo.Id,
		UserId:    bidMongo.UserId,
		AuctionId: bidMongo.AuctionId,
		Amount:    bidMongo.Amount,
		Timestamp: time.Unix(bidMongo.Timestamp, 0),
	}, nil
}
