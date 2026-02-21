package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()

	// Initialize the in-memory repository
	inmemRepo := repository.NewInmemRepository()

	// Initialize the trip service with the repository
	service := service.NewService(inmemRepo)

	// Example usage of the service
	// This would typically be called from an HTTP handler or another part of the application
	// For example:
	fare := &domain.RideFareModel{
		UserID:             primitive.NewObjectID(),
		PackageSlug:        "standard",
		TTotalPriceInCents: 1000,
	}

	trip, err := service.CreateTrip(ctx, fare)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created trip: %+v", trip)

	// Keep the application running to allow for further interactions
	for {
		time.Sleep(time.Second)
	}
}
