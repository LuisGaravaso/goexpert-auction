package main

import (
	"context"
	"log"

	"github.com/LuisGaravaso/goexpert-auction/configs/database/mongodb"
	auction_controller "github.com/LuisGaravaso/goexpert-auction/internal/infra/api/web/controller/auction"
	bid_controller "github.com/LuisGaravaso/goexpert-auction/internal/infra/api/web/controller/bid"
	user_controller "github.com/LuisGaravaso/goexpert-auction/internal/infra/api/web/controller/user"
	auction_repo "github.com/LuisGaravaso/goexpert-auction/internal/infra/database/auction"
	bid_repo "github.com/LuisGaravaso/goexpert-auction/internal/infra/database/bid"
	user_repo "github.com/LuisGaravaso/goexpert-auction/internal/infra/database/user"
	auction_usecase "github.com/LuisGaravaso/goexpert-auction/internal/usecase/auction"
	bid_usecase "github.com/LuisGaravaso/goexpert-auction/internal/usecase/bid"
	user_usecase "github.com/LuisGaravaso/goexpert-auction/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	ctx := context.Background()
	conn, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Start the server
	router := gin.Default()

	auctionController, bidController, userController := InitDependencies(conn)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auction/:auction_id", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auction/winner/:auction_id", auctionController.FindWinningBidInfo)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auction_id", bidController.FindBidByAuctionId)
	router.GET("/user/:user_id", userController.FindUserById)
	router.Run(":8080")
}

func InitDependencies(database *mongo.Database) (
	auctionController *auction_controller.AuctionController,
	bidController *bid_controller.BidController,
	userController *user_controller.UserController) {

	auctionRepository := auction_repo.NewAuctionRepository(database)
	bidRepository := bid_repo.NewBidRepository(database, auctionRepository)
	userRepository := user_repo.NewUserRepository(database)

	auctionUsecase := auction_usecase.NewAuctionUsecase(auctionRepository, bidRepository)
	bidUsecase := bid_usecase.NewBidUseCase(bidRepository)
	userUsecase := user_usecase.NewUserUsecase(userRepository)

	auctionController = auction_controller.NewAuctionController(auctionUsecase)
	bidController = bid_controller.NewBidController(bidUsecase)
	userController = user_controller.NewUserController(userUsecase)

	return
}
