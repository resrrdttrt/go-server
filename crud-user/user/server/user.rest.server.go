package server

import (
	"context"
	db "crud-user/user/database"
	endpoint "crud-user/user/endpoint"
	middleware "crud-user/user/middleware"
	model "crud-user/user/model"
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

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		schema.DecodeGetUserRequest,
		schema.EncodeResponse,
	))

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		schema.DecodeCreateUserRequest,
		schema.EncodeResponse,
	))

	r.Methods("PATCH").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateUser,
		schema.DecodeUpdateUserRequest,
		schema.EncodeResponse,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteUser,
		schema.DecodeDeleteUserRequest,
		schema.EncodeResponse,
	))

	return r
}

func StartUserHTTPServer(http_host string) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var httpAddr = flag.String("http", ":"+http_host, "http listen address")

	db := db.NewDB()
	flag.Parse()
	ctx := context.Background()
	var srv service.UserService
	{
		repository := model.NewRepo(db, logger)

		srv = service.NewService(repository, logger)
	}

	endpoints := endpoint.MakeEndpoints(srv)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}