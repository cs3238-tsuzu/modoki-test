package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "[docker image name]")
	}

	h2creds := credentials.NewTLS(&tls.Config{NextProtos: []string{"h2"}})
	conn, err := grpc.Dial(os.Getenv("MODOKI_API_SERVER"), grpc.WithTransportCredentials(h2creds))

	if err != nil {
		panic(err)
	}

	appClient := modoki.NewAppClient(conn)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", os.Getenv("MODOKI_API_SECRET"))

	resp, err := appClient.Deploy(ctx, &modoki.AppDeployRequest{
		Id: os.Getenv("MODOKI_APP_ID"),
		Spec: &modoki.AppSpec{
			Image: os.Args[1],
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
