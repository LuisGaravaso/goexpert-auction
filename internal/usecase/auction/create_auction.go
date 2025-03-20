package auction

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

func (u *AuctionUsecase) CreateAuction(ctx context.Context, input CreateAuctionInputDTO) (CreateAuctionOutputDTO, *internal_errors.InternalError) {
	auction, err := auction.NewAuction(
		input.ProductName,
		input.Category,
		input.Description,
		auction.ProductCondition(input.Condition),
	)

	if err != nil {
		return CreateAuctionOutputDTO{}, err
	}

	createdAuction, err := u.auctionRepository.CreateAuction(ctx, *auction)
	if err != nil {
		return CreateAuctionOutputDTO{}, err
	}

	return CreateAuctionOutputDTO{
		Id:          createdAuction.Id,
		ProductName: createdAuction.ProductName,
		Category:    createdAuction.Category,
		Description: createdAuction.Description,
		Condition:   ProductCondition(createdAuction.Condition),
		Status:      AuctionStatus(createdAuction.Status),
		Timestamp:   createdAuction.Timestamp,
	}, nil
}
