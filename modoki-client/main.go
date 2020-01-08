package main

import (
	"context"
	"fmt"
	"os"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"google.golang.org/grpc/metadata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "[docker image name]")
	}

	dialer := grpcutil.NewGRPCDialer("")
	dialer.StreamClientInterceptors = nil
	dialer.UnaryClientInterceptors = nil

	conn, err := dialer.Dial(os.Getenv("MODOKI_API_SERVER"), false)

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
