package bid

import "time"

// CreateUseCase DTOs
type CreateBidInputDTO struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type CreateBidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

// FindUseCase DTOs
type BidsFound struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type FindBidByAuctionIdInputDTO struct {
	AuctionId string `json:"auction_id"`
}

type FindBidByAuctionIdOutputDTO struct {
	Bids []BidsFound `json:"bids"`
}

// FindWinningBid DTOs
type FindWinningBidInputDTO struct {
	AuctionId string `json:"auction_id"`
}

type FindWinningBidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}
