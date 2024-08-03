package main

import (
    "net"

    "github.com/dawitel/addispay-project-2/configs"
    "github.com/dawitel/addispay-project-2/internal/order"
    "github.com/dawitel/addispay-project-2/internal/utils"
    "github.com/dawitel/addispay-project-2/internal/db"
    "github.com/dawitel/addispay-project-2/api/orders/pb"

    "google.golang.org/grpc"
)

var logger = utils.GetLogger()

func main() {

    config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Failed to load config: ", err)
    }

    if err := db.InitDB(); err != nil {
        logger.Error("Failed to initialize connection to database: ", err)
    }

    order.InitPulsar()

    lis, err := net.Listen("tcp", config.GRPCPort)
    if err != nil {
        logger.Error("Failed to listen: ", err)
    }

    s := grpc.NewServer()
    orderService := &order.Service{}
    pb.RegisterOrderServiceServer(s, orderService)

    logger.Success("Order service is running on port: ", config.GRPCPort)
    if err := s.Serve(lis); err != nil {
        logger.Error("Failed to serve: ", err)
    }
    
    // TODO: activate the publish order func
}
