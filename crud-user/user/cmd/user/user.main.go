package main

import (
	server "crud-user/user/server"
	utils "crud-user/user/utils"

	_ "github.com/lib/pq"
)

func main() {
	grpcHost := utils.GetEnv("GRPC_HOST", ":2222")
	httpHost := utils.GetEnv("HTTP_HOST", "1111")
	server.StartUserGRPCServer(grpcHost)
	server.StartUserHTTPServer(httpHost)
}
