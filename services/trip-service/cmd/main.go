package main

import (
	"fmt"
	"log"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/types"
	"time"
)

func main() {
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	//mux := http.NewServeMux()

	//httpHandler := h.HttpHandler{
	//	Service: svc,
	//}
	//
	//mux.HandleFunc("POST /preview", httpHandler.HandlePreviewTrip)
	//
	//server := &http.Server{
	//	Addr:    ":8083",
	//	Handler: mux,
	//}

	log.Println("Starting Trip Service running on port 8083")

	//if err := server.ListenAndServe(); err != nil {
	//	panic(err)
	//}

	t, err := svc.GetRoute(nil, &types.Coordinate{
		Latitude:  3.152581470137364,
		Longitude: 101.69921875,
	}, &types.Coordinate{
		Latitude:  3.152581470137364,
		Longitude: 101.70022583007812,
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", t)

	time.Sleep(1 * time.Second)

}
