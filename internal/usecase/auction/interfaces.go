package auction

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/entity/bid"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

type AuctionUsecase struct {
	auctionRepository auction.AuctionRepositoryInterface
	bidRepository     bid.BidRepositoryInterface
}

type AuctionUsecaseInterface interface {
	CreateAuction(ctx context.Context, input CreateAuctionInputDTO) (CreateAuctionOutputDTO, *internal_errors.InternalError)
	FindAuctionById(ctx context.Context, input FindAuctionByIdInputDTO) (FindAuctionByIdOutputDTO, *internal_errors.InternalError)
	FindAuctions(ctx context.Context, input FindAuctionsInputDTO) (FindAuctionsOutputDTO, *internal_errors.InternalError)
	FindWinningBidInfo(ctx context.Context, input WinningBidInfoInputDTO) (*WinningBidInfoOutputDTO, *internal_errors.InternalError)
}

func NewAuctionUsecase(auctionRepository auction.AuctionRepositoryInterface, bidRepository bid.BidRepositoryInterface) *AuctionUsecase {
	return &AuctionUsecase{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
	}
}
