package auction

import (
	"context"
	"net/http"

	"github.com/LuisGaravaso/goexpert-auction/configs/rest_err"
	"github.com/LuisGaravaso/goexpert-auction/internal/infra/api/web/validation"
	"github.com/LuisGaravaso/goexpert-auction/internal/usecase/auction"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuctionController struct {
	usecase auction.AuctionUsecaseInterface
}

func NewAuctionController(usecase auction.AuctionUsecaseInterface) *AuctionController {
	return &AuctionController{usecase: usecase}
}

func (a *AuctionController) CreateAuction(c *gin.Context) {
	var input auction.CreateAuctionInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx := context.Background()
	output, err := a.usecase.CreateAuction(ctx, input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (a *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auction_id")

	if err := uuid.Validate(auctionId); err != nil {
		restErr := rest_err.NewBadRequestError("invalid fields", rest_err.Causes{
			Field:   "auction_id",
			Message: "auction id must be a valid UUID",
		})
		c.JSON(restErr.Code, restErr)
		return
	}

	input := auction.FindAuctionByIdInputDTO{Id: auctionId}
	output, err := a.usecase.FindAuctionById(context.Background(), input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (a *AuctionController) FindAuctions(c *gin.Context) {
	var input auction.FindAuctionsInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		errRest := validation.ValidateErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx := context.Background()
	output, err := a.usecase.FindAuctions(ctx, input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (a *AuctionController) FindWinningBidInfo(c *gin.Context) {
	auctionId := c.Param("auction_id")

	if err := uuid.Validate(auctionId); err != nil {
		restErr := rest_err.NewBadRequestError("invalid fields", rest_err.Causes{
			Field:   "auction_id",
			Message: "auction id must be a valid UUID",
		})
		c.JSON(restErr.Code, restErr)
		return
	}

	input := auction.WinningBidInfoInputDTO{AuctionId: auctionId}
	output, err := a.usecase.FindWinningBidInfo(context.Background(), input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, output)
}
