package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"net/http"

	e "github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/utils/jsonresponse"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type (
	userServer struct {
		login *httptransport.Server
	}
)

func (us *userServer) GetLoginServer() *httptransport.Server {
	return us.login
}

func NewUserServer(set *e.EntrySet, logger log.Logger) *userServer {
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			return context.WithValue(ctx, "somekey", "somevalue")
		}),
		httptransport.ServerErrorHandler(newLogErrorHandler(logger)),
	}
	return &userServer{
		login: httptransport.NewServer(set.LoginEndPoint, DecodeUserLoginRequest, EncodeUserLoginResponse, options...),
	}
}

func DecodeUserLoginRequest(c context.Context, request *http.Request) (interface{}, error) {
	var req service.LoginRequest
	err := httpx.Parse(request, &req)
	if err != nil {
		return nil, errors.Wrap(err, "请求参数错误")
	}
	return &req, nil
}

func EncodeUserLoginResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(jsonresponse.ErrorResponse{Code: 2001, Message: f.Failed().Error()})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(jsonresponse.Response{Code: 0, Message: "success", Data: response})
}
