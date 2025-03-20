package bid

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/entity/bid"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

var bidBatch []bid.Bid

func (u *BidUseCase) Create(ctx context.Context, input CreateBidInputDTO) *internal_errors.InternalError {

	bid, err := bid.NewBid(input.UserId, input.AuctionId, input.Amount)
	if err != nil {
		return err
	}

	u.bidChannel <- *bid

	return nil

}
