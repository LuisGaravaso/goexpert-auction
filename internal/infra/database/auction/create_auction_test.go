package auction_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	auctionEntity "github.com/LuisGaravaso/goexpert-auction/internal/entity/auction"
	"github.com/LuisGaravaso/goexpert-auction/internal/infra/database/auction"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateAuction_ClosesAfterInterval(t *testing.T) {
	os.Setenv("AUCTION_INTERVAL", "2s") // tempo reduzido para teste

	log.Println(os.Getenv("AUCTION_INTERVAL"))

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		t.Fatalf("error connecting to mongodb: %v", err)
	}

	database := client.Database("auction_test")
	repo := auction.NewAuctionRepository(database)
	auc, _ := auctionEntity.NewAuction("product", "category", "description", auctionEntity.New)
	repo.CreateAuction(context.Background(), *auc)

	// wait for the auction to be closed
	time.Sleep(3 * time.Second)

	auc, err = repo.FindAuctionById(context.Background(), auc.Id)
	assert.Nil(t, err)
	assert.Equal(t, auctionEntity.Completed, auc.Status)
}
