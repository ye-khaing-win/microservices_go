package grpc_clients

import (
	"google.golang.org/grpc"
	"os"
	pb "ride-sharing/shared/proto/trip"
)

type TripServiceClient struct {
	Client pb.TripServiceClient
	conn   *grpc.ClientConn
}

func NewTripServiceClient() (*TripServiceClient, error) {
	tripServiceURL := os.Getenv("TRIP_SERVICE_URL")
	if tripServiceURL == "" {
		tripServiceURL = "trip-service:9093"
	}
	conn, err := grpc.NewClient(tripServiceURL)
	if err != nil {
		return nil, err
	}

	client := pb.NewTripServiceClient(conn)

	return &TripServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *TripServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
