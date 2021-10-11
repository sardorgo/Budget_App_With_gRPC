package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/satori/uuid"

	"google.golang.org/grpc"

	pb "github.com/sardorgo/app/proto/cost_proto"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "sardor"
	PASSWORD = "sardor"
	DBNAME   = "grpc_hw1"
)

type server struct {
	conn *sql.DB
	pb.UnimplementedCostServer
}

func (connection *server) CreateCost(ctx context.Context, req *pb.CreateCostRequest) (*pb.CostProfile, error) {
	db := connection.conn
	id := uuid.NewV4()
	req.CostProfile.Id = id.String()

	amount := req.GetCostProfile().GetAmount()
	summary := req.GetCostProfile().GetSummary()
	user_id := req.GetCostProfile().GetUserId()

	sqlInsert := `insert into cost (cost_id, cost_amount, summary, user_id) values ($1, $2, $3, $4)`

	if _, err := db.Exec(sqlInsert, id, int(amount), summary, user_id); err != nil {
		return nil, errors.Wrapf(err, "Cost Couldn't Be Inserted")
	}

	return req.CostProfile, nil
}

func main() {
	fmt.Println("Welcome to the server")
	lis, err := net.Listen("tcp", ":9300")

	if err != nil {
		errors.Wrap(err, "UserProfile couldn't be returned")
	}

	s := grpc.NewServer()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		errors.Wrap(err, "UserProfile couldn't be returned")
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		errors.Wrap(err, "User couldn't be listed")
	}

	pb.RegisterCostServer(s, &server{conn: db})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		errors.Wrap(err, "UserProfile couldn't be returned")
	}

}
