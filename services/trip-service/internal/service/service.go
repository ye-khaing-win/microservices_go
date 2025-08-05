package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ride-sharing/services/trip-service/internal/domain"
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
