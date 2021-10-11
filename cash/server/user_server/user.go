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

	pb "github.com/sardorgo/app/proto/user_proto"
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
	pb.UnimplementedUserProfilesServer
}

func (connection *server) CreateUser(ctx context.Context, req *pb.CreateUserProfileRequest) (*pb.UserProfile, error) {
	db := connection.conn
	id := uuid.NewV4()
	req.UserProfile.Id = id.String()

	firstName := req.GetUserProfile().GetFirstName()
	lastName := req.GetUserProfile().GetLastName()

	sqlInsert := `insert into users (user_id, first_name, last_name) values ($1, $2, $3);`

	if _, err := db.Exec(sqlInsert, id, firstName, lastName); err != nil {
		return nil, errors.Wrapf(err, "User couldn't be inserted")
	}

	return req.UserProfile, nil

}

func (connection *server) ListUsers(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUsersResponse, error) {
	db := connection.conn
	sqlStatement := `select user_id, first_name, last_name, total_money from users`
	result, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	defer result.Close()

	res := []*pb.UserProfile{}
	for result.Next() {
		var first, last, id string
		var total_money int
		if err = result.Scan(&id, &first, &last, &total_money); err != nil {
			errors.Wrap(err, "Users couln't be listed")
		}
		u := pb.UserProfile{
			Id:         id,
			FirstName:  first,
			LastName:   last,
			TotalMoney: int64(total_money),
		}
		res = append(res, &u)
	}
	ans := pb.ListUsersResponse{Profiles: res}
	return &ans, nil
}

func main() {
	fmt.Println("Welcome to the server")
	lis, err := net.Listen("tcp", ":9500")

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

	pb.RegisterUserProfilesServer(s, &server{conn: db})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		errors.Wrap(err, "UserProfile couldn't be returned")
	}

}
