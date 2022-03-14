package transport

import (
	"context"
	"errors"
	"fmt"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/cmd/grpc/proto/pb"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/common/log/logx"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
)

type (
	logErrorHandler struct {
		logger logx.Logger
	}

	userGrpcServer struct {
		login grpctransport.Handler
	}
)

func NewUserGrpcServer(set *endpoint.EntrySet) pb.UserServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			return context.WithValue(ctx, "somekey", "somevalue")
		}),
	}
	return &userGrpcServer{
		login: grpctransport.NewServer(set.UserEndPoint.LoginEndPoint, DecodeGrpcUserLoginRequest, EncodeGrpcUserLoginResponse, options...),
	}
}

func (s *userGrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, res, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginReply), nil
}

func (s *userGrpcServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	_, res, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LogoutReply), nil
}

func (s *userGrpcServer) Info(ctx context.Context, req *pb.InfoRequest) (*pb.InfoReply, error) {
	_, res, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InfoReply), nil
}

func newLogErrorHandler(logger logx.Logger) *logErrorHandler {
	return &logErrorHandler{
		logger: logger,
	}
}

func (h *logErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Errorf("err: %s", err)
}

func DecodeGrpcUserLoginRequest(ctx context.Context, request interface{}) (interface{}, error) {
	loginRequest, ok := request.(*pb.LoginRequest)
	if !ok {
		return nil, errors.New("invalid request type")
	}
	return &service.LoginRequest{
		Username: loginRequest.GetUsername(),
		Password: loginRequest.GetPassword(),
	}, nil
}

func EncodeGrpcUserLoginResponse(ctx context.Context, response interface{}) (interface{}, error) {
	fmt.Println(response)
	fmt.Printf("%T\n", response)
	reply, ok := response.(*service.LoginResponse)
	if !ok {
		return nil, errors.New("invalid response type1")
	}
	return &pb.LoginReply{Token: reply.Token, ExpireSeconds: int64(reply.ExpireSeconds)}, nil
}
