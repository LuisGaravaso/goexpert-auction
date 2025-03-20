package auction

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"github.com/LuisGaravaso/goexpert-auction/internal/usecase/bid"
)

func (u *AuctionUsecase) FindAuctionById(ctx context.Context, input FindAuctionByIdInputDTO) (FindAuctionByIdOutputDTO, *internal_errors.InternalError) {
	auction, err := u.auctionRepository.FindAuctionById(ctx, input.Id)
	if err != nil {
		return FindAuctionByIdOutputDTO{}, err
	}

	return FindAuctionByIdOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}, nil
}

func (u *AuctionUsecase) FindAuctions(ctx context.Context, input FindAuctionsInputDTO) (FindAuctionsOutputDTO, *internal_errors.InternalError) {
	auctions, err := u.auctionRepository.FindAuctions(ctx, auction.AuctionStatus(input.Status), input.Category, input.ProductName)
	if err != nil {
		return FindAuctionsOutputDTO{}, err
	}

	var auctionsDTO []FindAuctionByIdOutputDTO
	for _, auction := range auctions {
		auctionsDTO = append(auctionsDTO, FindAuctionByIdOutputDTO{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}

	return FindAuctionsOutputDTO{Auctions: auctionsDTO}, nil
}

func (u *AuctionUsecase) FindWinningBidInfo(ctx context.Context, input WinningBidInfoInputDTO) (*WinningBidInfoOutputDTO, *internal_errors.InternalError) {
	auction, err := u.auctionRepository.FindAuctionById(ctx, input.AuctionId)
	if err != nil {
		return nil, err
	}
	auctionDto := FindAuctionByIdOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	winningBid, err := u.bidRepository.FindWinningBidByAuctionId(ctx, input.AuctionId)
	if err != nil {
		return &WinningBidInfoOutputDTO{
			Auction: auctionDto,
			Bid:     nil,
		}, nil
	}

	return &WinningBidInfoOutputDTO{
		Auction: auctionDto,
		Bid: &bid.FindWinningBidOutputDTO{
			Id:        winningBid.Id,
			UserId:    winningBid.UserId,
			AuctionId: winningBid.AuctionId,
			Amount:    winningBid.Amount,
			Timestamp: winningBid.Timestamp,
		},
	}, nil
}
