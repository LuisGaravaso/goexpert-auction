package bid

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/configs/logger"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/bid"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

type BidUseCase struct {
	bidRepository bid.BidRepositoryInterface

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid.Bid
}

func NewBidUseCase(bidRepository bid.BidRepositoryInterface) *BidUseCase {

	maxSizeInteval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()
	bidUseCase := &BidUseCase{
		bidRepository:       bidRepository,
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: maxSizeInteval,
		timer:               time.NewTimer(maxSizeInteval),
		bidChannel:          make(chan bid.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

type BidUseCaseInterface interface {
	Create(ctx context.Context, input CreateBidInputDTO) *internal_errors.InternalError
	FindBidByAuctionId(ctx context.Context, input FindBidByAuctionIdInputDTO) (*FindBidByAuctionIdOutputDTO, *internal_errors.InternalError)
	FindWinningBid(ctx context.Context, input FindWinningBidInputDTO) (*FindWinningBidOutputDTO, *internal_errors.InternalError)
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return time.Minute * 3
	}

	return duration
}

func getMaxBatchSize() int {
	maxBatchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}

	return maxBatchSize
}

func (u *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(u.bidChannel)

		for {
			select {
			case bid, ok := <-u.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := u.bidRepository.Create(ctx, bidBatch); err != nil {
							logger.Error("Error trying to process bid batch list", err)
							return
						}
					}
				}

				bidBatch = append(bidBatch, bid)
				if len(bidBatch) >= u.maxBatchSize {
					if err := u.bidRepository.Create(ctx, bidBatch); err != nil {
						logger.Error("Error trying to process bid batch list", err)
						return
					}
					bidBatch = nil
					u.timer.Reset(u.batchInsertInterval)
				}

			case <-u.timer.C:
				if len(bidBatch) > 0 {
					if err := u.bidRepository.Create(ctx, bidBatch); err != nil {
						logger.Error("Error trying to process bid batch list", err)
						return
					}
				}
				bidBatch = nil
				u.timer.Reset(u.batchInsertInterval)

			}
		}
	}()
}
