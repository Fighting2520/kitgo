package main

import (
	"context"
	"fmt"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/cmd/grpc/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	login, err := client.Login(context.Background(), &pb.LoginRequest{
		Username: "jkwang",
		Password: "111111",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(login.Token)
	fmt.Println(login.ExpireSeconds)
}
