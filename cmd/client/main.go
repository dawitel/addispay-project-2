package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/dawitel/addispay-project-2/api/orders/pb" // Replace with the correct import path
	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger = utils.GetLogger()

func main() {
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Failed to load config: ", err)
    }
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/checkout", APIgatewayHandler)
	mux.HandleFunc("/processed-transactions", OrderResponseHandler)
	srv := &http.Server{
		Addr: config.APIGatewayAddr,
		Handler: mux,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// Start the HTTP server.
	logger.Success("starting API gateway server on: ",  srv.Addr)
	if err = srv.ListenAndServe(); err != nil {
		logger.Error("Faied to serve the API gateway: ", err)
	}

}

func APIgatewayHandler(w http.ResponseWriter, r *http.Request) {
	config, err := configs.LoadConfig()
    if err != nil {
		logger.Error("Failed to load configuration files: ", err)
    }

	// Set up a connection to the server.
	conn, err := grpc.NewClient(config.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to the order service: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewOrderServiceClient(conn)
	order := &pb.OrderRequest{}
	if err = json.NewDecoder(r.Body).Decode(order); err != nil {
		logger.Error("Failed to decode the request object: ", err)
	}

	// Contact the server and create a new order
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("Could not create order: %v", err)
	}
	log.Printf("Order Response: %s", resp.GetStatus())
}


func OrderResponseHandler(w http.ResponseWriter, r *http.Request) {
	config, err := configs.LoadConfig()
    if err != nil {
		logger.Error("Failed to load configuration files: ", err)
    }
    // Forward the data to the forntend client
    clientURL := fmt.Sprintf("%s/api/processed-transactions", config.FrontendAddr) // Assuming an endpoint in Next.js to receive data
    req, err := http.NewRequest("POST", clientURL, r.Body)
    if err != nil {
        http.Error(w, "Failed to create request to the Frontend", http.StatusInternalServerError)
		logger.Error("Failed to create request to the Frontend: ", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Failed to forward request", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Copy the response from the forntend client back to the order service processor
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
