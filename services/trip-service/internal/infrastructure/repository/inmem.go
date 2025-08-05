package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type InmemRepository struct {
	trips     map[string]*domain.Trip
	rideFares map[string]*domain.RideFare
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		trips:     make(map[string]*domain.Trip),
		rideFares: make(map[string]*domain.RideFare),
	}
}

func (r *InmemRepository) CreateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
