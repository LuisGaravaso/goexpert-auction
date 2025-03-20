package bid

import (
	"context"
	"net/http"

	"github.com/LuisGaravaso/goexpert-auction/configs/rest_err"
	"github.com/LuisGaravaso/goexpert-auction/internal/infra/api/web/validation"
	"github.com/LuisGaravaso/goexpert-auction/internal/usecase/bid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BidController struct {
	usecase bid.BidUseCaseInterface
}

func NewBidController(usecase bid.BidUseCaseInterface) *BidController {
	return &BidController{usecase: usecase}
}

func (b *BidController) CreateBid(c *gin.Context) {
	var input bid.CreateBidInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx := context.Background()
	err := b.usecase.Create(ctx, input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (b *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auction_id")

	if err := uuid.Validate(auctionId); err != nil {
		restErr := rest_err.NewBadRequestError("invalid fields", rest_err.Causes{
			Field:   "auction_id",
			Message: "auction id must be a valid UUID",
		})
		c.JSON(restErr.Code, restErr)
		return
	}

	input := bid.FindBidByAuctionIdInputDTO{AuctionId: auctionId}
	bidData, err := b.usecase.FindBidByAuctionId(context.Background(), input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidData)
}
