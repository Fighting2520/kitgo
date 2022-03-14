package main

import (
	"flag"
	"fmt"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/cmd/http/config"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/endpoint"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/model"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/transport"
	"github.com/Fighting2520/kitgo/common/configx/nacosconfig"
	"github.com/Fighting2520/kitgo/common/log/logx"
	"github.com/Fighting2520/kitgo/common/middleware"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"strconv"
)

var (
	configHost      = flag.String("configHost", "", "the config center host")
	configPort      = flag.Int("configPort", 0, "the config center port")
	configNamespace = flag.String("configNamespace", "", "the config center namespace")
	configDataId    = flag.String("configDataId", "", "the config center dataId")
	configGroup     = flag.String("configGroup", "", "the config center dataId")
)

func main() {
	flag.Parse()
	var conf config.Config
	err := unmarshalConfigFromNacos(nacosconfig.NewConfig(
		nacosconfig.WithIpAddr(*configHost),
		nacosconfig.WithPort(uint64(*configPort)),
		nacosconfig.WithNamespaceId(*configNamespace),
	), *configDataId, *configGroup, &conf)
	if err != nil {
		log.Fatal(err)
	}
	// 初始化日志
	logx.MustSetup(conf.Logs)
	userModel := model.NewUserModel(conf.UserMysql.Dsn, conf.UserMysql.Table.User)
	userService := service.NewService(userModel)
	var userEndPoint = endpoint.NewUserEndPoint(userService, rate.NewLimiter(rate.Limit(conf.RateLimit.Limit), conf.RateLimit.Burst))
	entrySet := endpoint.NewEntrySet(userEndPoint)
	userServer := transport.NewUserServer(entrySet)
	r := mux.NewRouter()
	r.Use(middleware.Recovery)
	r.Methods(http.MethodPost).Path("/login").Handler(userServer.GetLoginServer())
	server := http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(conf.Port),
	}
	server.ListenAndServe()
}

func unmarshalConfigFromNacos(c *nacosconfig.Config, dataId, group string, v interface{}) error {
	client, err := nacosconfig.NewClient(
		nacosconfig.WithIpAddr(c.IpAddr),
		nacosconfig.WithPort(c.Port),
		nacosconfig.WithNamespaceId(c.NamespaceId),
	)
	if err != nil {
		return err
	}
	configStr, err := client.GetConfig(dataId, group)
	if err != nil {
		return err
	}
	fmt.Println(configStr)
	err = yaml.Unmarshal([]byte(configStr), v)
	fmt.Println(v)
	return nil
}
