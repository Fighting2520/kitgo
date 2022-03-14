package main

import (
	"fmt"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/cmd/grpc/proto/pb"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/model"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/transport"
	"github.com/Fighting2520/kitgo/common/log/logx"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net"
)

func main() {
	userModel := model.NewUserModel("root:123456@tcp(127.0.0.1:3306)/semantic_map?charset=utf8&parseTime=True&loc=Asia%2FShanghai", "user")
	logx.MustSetup(logx.LogConf{
		Mode: "console",
	})
	userService := service.NewService(userModel)
	limit := rate.NewLimiter(1, 1)
	var userEndPoint = endpoint.NewUserEndPoint(userService, limit)
	entrySet := endpoint.NewEntrySet(userEndPoint)
	grpcServ := transport.NewUserGrpcServer(entrySet)
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
