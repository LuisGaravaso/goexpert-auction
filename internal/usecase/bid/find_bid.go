package bid

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

func (u *BidUseCase) FindBidByAuctionId(ctx context.Context, input FindBidByAuctionIdInputDTO) (*FindBidByAuctionIdOutputDTO, *internal_errors.InternalError) {
	bids, err := u.bidRepository.FindBidByAuctionId(ctx, input.AuctionId)
	if err != nil {
		return nil, err
	}

	var bidsOutputList FindBidByAuctionIdOutputDTO
	for _, bid := range bids {
		bidsOutputList.Bids = append(bidsOutputList.Bids, BidsFound{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}
	return &bidsOutputList, nil
}

func (u *BidUseCase) FindWinningBid(ctx context.Context, input FindWinningBidInputDTO) (*FindWinningBidOutputDTO, *internal_errors.InternalError) {
	bid, err := u.bidRepository.FindWinningBidByAuctionId(ctx, input.AuctionId)
	if err != nil {
		return nil, err
	}

	return &FindWinningBidOutputDTO{
		Id:        bid.Id,
		UserId:    bid.UserId,
		AuctionId: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}, nil
}
