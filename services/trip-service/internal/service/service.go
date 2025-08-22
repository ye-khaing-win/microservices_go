package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type Service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTrip(ctx context.Context, fare *domain.RideFare) (*domain.Trip, error) {
	t := &domain.Trip{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)
}

func (s *Service) GetRoute(ctx context.Context, pickup *types.Coordinate, destination *types.Coordinate) (*types.OSRMApiResponse, error) {
	log.Println("Service: ", destination.Latitude, destination.Longitude, pickup.Latitude, pickup.Longitude)

	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false",
		pickup.Longitude, pickup.Latitude,
		destination.Longitude, destination.Latitude)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get route from OSRM API: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var routeResp types.OSRMApiResponse
	if err := json.Unmarshal(body, &routeResp); err != nil {
		return nil, fmt.Errorf("failed to parse OSRM API response: %v", err)
	}

	return &routeResp, nil
}
