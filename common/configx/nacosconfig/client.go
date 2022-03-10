package nacosconfig

import (
	"errors"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type (
	Client struct {
		config_client.IConfigClient
	}
)

var ErrInvalidServerAddress = errors.New("invalid server address")

func NewClient(options ...Option) (*Client, error) {
	var conf Config
	for _, option := range options {
		option(&conf)
	}
	if conf.IpAddr == "" {
		return nil, ErrInvalidServerAddress
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ServerConfigs: []constant.ServerConfig{*constant.NewServerConfig(conf.IpAddr, conf.Port, convertServerOption(conf)...)},
		ClientConfig:  constant.NewClientConfig(convertClientOption(conf)...),
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		client,
	}, nil
}

func (c *Client) GetConfig(dataId, group string) (string, error) {
	return c.IConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}
