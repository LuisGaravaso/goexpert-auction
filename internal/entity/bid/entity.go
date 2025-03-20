package bid

import (
	"context"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"github.com/google/uuid"
)

type Bid struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type BidRepositoryInterface interface {
	Create(ctx context.Context, bidEntities []Bid) *internal_errors.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_errors.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_errors.InternalError)
}

func NewBid(userId, auctionId string, amount float64) (*Bid, *internal_errors.InternalError) {
	bid := &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_errors.InternalError {
	if err := uuid.Validate(b.UserId); err != nil {
		return internal_errors.NewBadRequestError("Invalid user id")
	}

	if err := uuid.Validate(b.AuctionId); err != nil {
		return internal_errors.NewBadRequestError("Invalid auction id")
	}

	if b.Amount <= 0 {
		return internal_errors.NewBadRequestError("Invalid amount")
	}

	return nil

}
