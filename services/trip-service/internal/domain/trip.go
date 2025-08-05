package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trip struct {
	ID       primitive.ObjectID
	UserID   string
	Status   string
	RideFare *RideFare
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *Trip) (*Trip, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFare) (*Trip, error)
}
