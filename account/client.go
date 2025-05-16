package account

import (
	"github.com/yash170603/goservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	// Replace deprecated grpc.WithInsecure() with the newer approach
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewAccountServiceClient(conn)
	return &Client{
		conn,
		client,
	}, nil
}
