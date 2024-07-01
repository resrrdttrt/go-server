package fmain

import (
	// "google.golang.org/grpc"

	"context"
	pb "crud-user/user/pb"
	"log"
	"time"

	server "crud-user/user/server"
)

func fmain() {
	// Set up a connection to the server.
	// conn, err := grpc.Dial("localhost:2222", grpc.WithInsecure(), grpc.WithBlock())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewUserServiceClient(conn)

	// // Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	host := "localhost:2222"
	grpcClient, err := server.StartUserGRPCClient(host)
	if err != nil {
		log.Fatalf("failed to start gRPC client: %v", err)
	}
	defer grpcClient.Close()

	client := grpcClient.GetClient()

	defer cancel()
	r, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
		User: &pb.User{
			Id:       "194fc53b-c3ff-4aaa-867e-348ddd490d47",
			Email:    "newemail@example.com",
			Password: "newpassword2222",
		},
	})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	// log.Printf("User: %s, Email: %s, Password: %s", r.User.Id, r.User.Email, r.User.Password)
	log.Printf("Success: %t", r.Success)
}
