package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"github.com/dawitel/addispay-project-2/api/orders/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// var orderServiceAddr string

// func main() {
// 	// Set up a connection to the order server.
//     fmt.Println("Connecting to order service via", orderServiceAddr)
// 	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("could not connect to order service: %v", err)
// 	}
// 	defer conn.Close()
// j
// 	// Register gRPC server endpoint
// 	// Note: Make sure the gRPC server is running properly and accessible
// 	mux := runtime.NewServeMux()
// 	if err = pb.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
// 		log.Fatalf("failed to register the order server: %v", err)
// 	}

// 	// start listening to requests from the gateway server
// 	addr := "0.0.0.0:8080"
// 	fmt.Println("API gateway server is running on " + addr)
// 	if err = http.ListenAndServe(addr, mux); err != nil {
// 		log.Fatal("gateway server closed abruptly: ", err)
// 	}
// }

// // package main

// // import (
// // 	"context"
// // 	"log"
// // 	"time"

// // 	"google.golang.org/grpc"
// // 	pb "path_to_your_project/proto" // Replace with the correct import path
// // )

// // const (
// // 	address = "localhost:50051" // gRPC server address
// // )

// // func main() {
// // 	// Set up a connection to the server.
// // 	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
// // 	if err != nil {
// // 		log.Fatalf("Did not connect: %v", err)
// // 	}
// // 	defer conn.Close()
// // 	client := pb.NewOrderServiceClient(conn)

// // 	// Contact the server and create a new order
// // 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// // 	defer cancel()

// // 	order := &pb.OrderRequest{
// // 		OrderId:    "12345",
// // 		Amount:     100.50,
// // 		PhoneNumber: "123-456-7890",
// // 		CallbackUrl: "http://example.com/callback",
// // 	}
// // 	resp, err := client.CreateOrder(ctx, order)
// // 	if err != nil {
// // 		log.Fatalf("Could not create order: %v", err)
// // 	}
// // 	log.Printf("Order Response: %s", resp.GetStatus())
// // }

