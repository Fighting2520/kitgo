package main

import (
	"net/http"
	"os"

	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/model"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/transport"
	"github.com/Fighting2520/kitgo/common/middleware"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func main() {
	r := mux.NewRouter()
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
	userServer := transport.NewUserServer(entrySet, logger)
	r.Use(middleware.Recovery)
	r.Path("/demo").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		panic("this is panic demo")
	})
	r.Methods(http.MethodPost).Path("/login/{id}").Handler(userServer.GetLoginServer())
	server := http.Server{
		Handler: r,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
