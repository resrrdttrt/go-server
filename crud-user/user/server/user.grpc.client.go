package server

import (
	"log"
	"time"

	pb "crud-user/user/pb"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	client pb.UserServiceClient
	conn   *grpc.ClientConn
}

// StartUserGRPCClient establishes a connection to the gRPC server and returns a GRPCClient.
func StartUserGRPCClient(host string) (*GRPCClient, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)

	return &GRPCClient{
		client: client,
		conn:   conn,
	}, nil
}

// GetClient returns the UserServiceClient.
func (c *GRPCClient) GetClient() pb.UserServiceClient {
	return c.client
}

// Close gracefully closes the connection to the gRPC server.
func (c *GRPCClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Fatalf("failed to close connection: %v", err)
	}
}
