package main

import (
	server "crud-user/user/server"
	utils "crud-user/user/utils"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	grpcHost := utils.GetEnv("GRPC_HOST", ":2222")
	grpcClient, err := server.StartUserGRPCClient(grpcHost)
	if err != nil {
		log.Fatalf("failed to start gRPC client: %v", err)
	}
	defer grpcClient.Close()

	grpc_client := grpcClient.GetClient()
	httpHost := utils.GetEnv("HTTP_HOST", "3333")
	server.StartGatewayHTTPServer(httpHost, grpc_client)
}
