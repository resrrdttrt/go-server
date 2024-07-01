package server

import (
	"context"
	db "crud-user/user/database"
	pb "crud-user/user/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type myUserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s myUserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := req.GetUserId()
	db_conn := db.NewDB()
	var user_email string
	var password string
	err := db_conn.QueryRow("SELECT email, password FROM users WHERE id=$1", id).Scan(&user_email, &password)
	if err != nil {
		log.Printf("User not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	return &pb.GetUserResponse{User: &pb.User{Id: id, Email: user_email, Password: password}}, nil
}

func (s myUserServiceServer) UpdateUser(ctx context.Context,req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := req.GetUser()
	db_conn := db.NewDB()
	var user_email string
	err := db_conn.QueryRow("SELECT email FROM users WHERE id=$1", user.Id).Scan(&user_email)
	if err != nil {
		log.Printf("User not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Update the user's email in the database
	_, err = db_conn.Exec("UPDATE users SET email=$1, password=$2 WHERE id=$3", user.GetEmail(), user.GetPassword(), user.Id)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update user")
	}

	// Return the updated user
	return &pb.UpdateUserResponse{Success: true}, nil

}

func StartUserGRPCServer(host string) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myUserServiceServer{}
	pb.RegisterUserServiceServer(serverRegistrar, service)

	go func() {
		err = serverRegistrar.Serve(lis)
		if err != nil {
			log.Fatalf("impossible to serve: %s", err)
		}
	}()
}
