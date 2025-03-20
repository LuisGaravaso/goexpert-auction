package bid

import (
	"context"
	"sync"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/bid"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

func (r *BidRepository) Create(ctx context.Context, bidEntities []bid.Bid) *internal_errors.InternalError {
	var wg sync.WaitGroup
	for _, bidEntity := range bidEntities {
		wg.Add(1)
		go func(bidEntity bid.Bid) {
			defer wg.Done()
			auctionEntity, err := r.AuctionRepository.FindAuctionById(ctx, bidEntity.AuctionId)
			if err != nil {
				logger.Error("Error finding auction by id", err)
				return
			}

			if auctionEntity.Status != auction.Active {
				return
			}

			bidMongo := &BidMongo{
				Id:        bidEntity.Id,
				UserId:    bidEntity.UserId,
				AuctionId: bidEntity.AuctionId,
				Amount:    bidEntity.Amount,
				Timestamp: bidEntity.Timestamp.Unix(),
			}

			if _, err := r.Collection.InsertOne(ctx, bidMongo); err != nil {
				logger.Error("Error inserting bid", err)
				return
			}

		}(bidEntity)
	}

	wg.Wait()
	return nil
}
