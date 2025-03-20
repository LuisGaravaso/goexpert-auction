package auction

import (
	"context"
	"time"

	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
	"github.com/google/uuid"
)

type Auction struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp"`
}

type ProductCondition int
type AuctionStatus int

const (
	Active    AuctionStatus = iota // 0
	Completed                      // 1
)

const (
	New         ProductCondition = iota // 0
	Used                                // 1
	Refurbished                         // 2
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction Auction) (*Auction, *internal_errors.InternalError)
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_errors.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category string, productName string) ([]Auction, *internal_errors.InternalError)
}

func NewAuction(productName, category, description string, condition ProductCondition) (*Auction, *internal_errors.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() *internal_errors.InternalError {
	if len(a.ProductName) <= 1 {
		return internal_errors.NewBadRequestError("product name must be at least 2 characters long")
	}
	if len(a.Category) <= 1 {
		return internal_errors.NewBadRequestError("category must be at least 2 characters long")
	}
	if len(a.Description) <= 1 {
		return internal_errors.NewBadRequestError("description must be at least 2 characters long")
	}
	if a.Condition != New && a.Condition != Refurbished && a.Condition != Used {
		return internal_errors.NewBadRequestError("invalid product condition")
	}
	return nil
}
