package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pbCash "github.com/sardorgo/app/proto/cash_proto"
	pbCost "github.com/sardorgo/app/proto/cost_proto"
	pbUser "github.com/sardorgo/app/proto/user_proto"
)

var userConnection pbUser.UserProfilesClient
var cashConnection pbCash.CashClient
var costConnection pbCost.CostClient

// Functions for users --> user_server - [users]

func PostUser(c *gin.Context) {
	var newUser pbUser.UserProfile

	if err := c.ShouldBindJSON(&newUser); err != nil {
		return
	}

	req := &pbUser.CreateUserProfileRequest{
		UserProfile: &pbUser.UserProfile{
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Id:        newUser.Id,
		},
	}

	res, err := userConnection.CreateUser(context.Background(), req)

	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, res)
}

func ListAll(c *gin.Context) {
	req := &pbUser.ListUserRequest{}
	res, err := userConnection.ListUsers(context.Background(), req)

	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

// Function for cash --> cash_server -- [cash]

func PostCash(c *gin.Context) {
	var newCash pbCash.CashProfile

	if err := c.ShouldBindJSON(&newCash); err != nil {
		return
	}

	req := &pbCash.CreateCashRequest{
		CashProfile: &pbCash.CashProfile{
			Amount:  newCash.Amount,
			Summary: newCash.Summary,
			UserId:  newCash.UserId,
			Id:      newCash.Id,
		},
	}

	res, err := cashConnection.CreateCash(context.Background(), req)

	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

// Function for cost --> cost_server -- [cost]

func PostCost(c *gin.Context) {
	var newCost pbCost.CostProfile

	if err := c.ShouldBindJSON(&newCost); err != nil {
		return
	}

	req := &pbCost.CreateCostRequest{
		CostProfile: &pbCost.CostProfile{
			Amount:  newCost.Amount,
			Summary: newCost.Summary,
			UserId:  newCost.UserId,
			Id:      newCost.Id,
		},
	}

	res, err := costConnection.CreateCost(context.Background(), req)

	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	fmt.Println("Welcome Client")
	conn, err := grpc.Dial("localhost:9500", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer conn.Close()

	userConnection = pbUser.NewUserProfilesClient(conn)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	// ----------------------------------------------

	fmt.Println("Welcome Client")
	conn1, err := grpc.Dial("localhost:9400", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer conn.Close()

	cashConnection = pbCash.NewCashClient(conn1)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	// -----------------------------------------------

	fmt.Println("Welcome Client")
	conn2, err := grpc.Dial("localhost:9300", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer conn.Close()

	costConnection = pbCost.NewCostClient(conn2)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	router := gin.Default()

	router.POST("/users", PostUser)
	router.POST("/cost", PostCost)
	router.POST("/cash", PostCash)

	router.GET("/users", ListAll)

	router.Run("localhost:5000")
}
