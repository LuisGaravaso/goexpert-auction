package auction

import (
	"time"

	"github.com/LuisGaravaso/goexpert-auction/internal/usecase/bid"
)

type CreateAuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=1"`
	Description string           `json:"description" binding:"required,min=1,max=200"`
	Condition   ProductCondition `json:"condition"`
}

type CreateAuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int64
type AuctionStatus int64

type FindAuctionByIdInputDTO struct {
	Id string `json:"id"`
}

type FindAuctionByIdOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type FindAuctionsInputDTO struct {
	Status      AuctionStatus `json:"status"`
	Category    string        `json:"category"`
	ProductName string        `json:"product_name"`
}

type FindAuctionsOutputDTO struct {
	Auctions []FindAuctionByIdOutputDTO `json:"auctions"`
}

type WinningBidInfoInputDTO struct {
	AuctionId string `json:"auction_id"`
}

type WinningBidInfoOutputDTO struct {
	Auction FindAuctionByIdOutputDTO     `json:"auction"`
	Bid     *bid.FindWinningBidOutputDTO `json:"bid,omitempty"`
}
