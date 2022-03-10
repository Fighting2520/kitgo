package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/cmd/grpc/proto/pb"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/model"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/transport"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

func main() {
	userModel := model.NewUserModel("root:123456@tcp(127.0.0.1:3306)/semantic_map?charset=utf8&parseTime=True&loc=Asia%2FShanghai", "user")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	userService := service.Chain(service.LoggingMiddleware(logger))(service.NewService(userModel))
	limit := rate.NewLimiter(1, 1)
	var userEndPoint = endpoint.NewUserEndPoint(userService, logger, limit)
	entrySet := endpoint.NewEntrySet(userEndPoint)
	grpcServ := transport.NewUserGrpcServer(entrySet, logger)
	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterUserServer(srv, grpcServ)
	if err = srv.Serve(listen); err != nil {
		fmt.Println(err)
		return
	}

}
