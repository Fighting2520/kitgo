package transport

import (
	"context"
	"encoding/json"
	"github.com/Fighting2520/kitgo/common/errorx"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	e "github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/utils/jsonresponse"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/rest/httpx"
)

type (
	userServer struct {
		login *httptransport.Server
	}
)

func (us *userServer) GetLoginServer() *httptransport.Server {
	return us.login
}

func NewUserServer(set *e.EntrySet) *userServer {
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			return context.WithValue(ctx, "somekey", "somevalue")
		}),
		httptransport.ServerErrorHandler(newDefaultErrorHandler()),
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
			w.Header().Set("Content-Type", contentType)
			code := http.StatusInternalServerError
			if _, ok := err.(*errorx.CodeError); ok {
				code = http.StatusNotAcceptable
			}
			w.WriteHeader(code)
			w.Write(body)
		}),
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
		err := errorx.NewCodeError(2001, f.Failed().Error())
		return json.NewEncoder(w).Encode(jsonresponse.ErrorResponse{Code: err.GetCode(), Message: err.GetMessage()})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(jsonresponse.Response{Code: 0, Message: "success", Data: response})
}
