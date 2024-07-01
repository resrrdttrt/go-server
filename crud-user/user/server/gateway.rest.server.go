package server

import (
	"context"
	endpoint "crud-user/user/endpoint"
	middleware "crud-user/user/middleware"
	pb "crud-user/user/pb"
	schema "crud-user/user/schema"
	service "crud-user/user/service"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func NewGatewayHTTPServer(ctx context.Context, gwendpoints endpoint.GWEndpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)

	r.Methods("GET").Path("/gateway/{id}").Handler(httptransport.NewServer(
		gwendpoints.GetUser,
		schema.DecodeGetUserRequest,
		schema.EncodeResponse,
	))

	r.Methods("PATCH").Path("/gateway/{id}").Handler(httptransport.NewServer(
		gwendpoints.UpdateUser,
		schema.DecodeUpdateUserRequest,
		schema.EncodeResponse,
	))

	return r
}

func StartGatewayHTTPServer(http_host string, grpc_client pb.UserServiceClient) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "gateway",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var httpAddr = flag.String("http", ":"+http_host, "http listen address")
	flag.Parse()
	ctx := context.Background()
	var srv service.GatewayService

	{

		srv = service.NewGatewayService(grpc_client, logger)
	}

	endpoints := endpoint.MakeGWEndpoints(srv)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := NewGatewayHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
