package nacosconfig

import "testing"

var client *Client

func TestNewClient(t *testing.T) {
	var err error
	client, err = NewClient(
		WithIpAddr("192.168.2.231"),
		WithPort(8848),
		WithNamespaceId("4363bd87-070b-4961-ab88-b1610c8295c0"),
		WithLogDir("/home/yelai/golangProjects/demo/github.com/Fighting2520/kitgo/modules/configx/nacosconfig/log"),
		WithTimeoutMs(3000),
		WithCacheDir("/home/yelai/golangProjects/demo/github.com/Fighting2520/kitgo/modules/configx/nacosconfig/cache"),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetConfig(t *testing.T) {
	TestNewClient(t)
	config, err := client.GetConfig("config", "local")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
}
