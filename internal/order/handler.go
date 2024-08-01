package order

import (
    "context"
    "github.com/dawitel/addispay-project-2/api/orders/pb"
    "github.com/dawitel/addispay-project-2/internal/models"
    "github.com/dawitel/addispay-project-2/internal/utils"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

var logger = utils.GetLogger()

// Service implements the OrderServiceServer interface defined in order.proto.
type Service struct {
    pb.UnimplementedOrderServiceServer
}

// CreateOrder handles the creation of an order and publishes it to a Pulsar topic.
func (s *Service) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    orderID := utils.GenerateID() 
    // Validate request
    if req.CustId == "" || req.Amount <= 0 {
        return nil, status.Error(codes.InvalidArgument, "Invalid order details")
    }

    // Convert gRPC request to internal model
    orderModel := &models.Order{
        OrderID:     orderID,
        CustID:      req.CustId,
        CustBankAcc: req.CustBankAcc,
        Amount:      req.Amount,
        PhoneNumber: req.PhoneNumber,
        CallbackURL: req.CallbackUrl,
    }

    // Store order in the database
    if err := SaveOrder(orderModel); err != nil {
        logger.Error("Failed to save order:", err)
        return nil, status.Error(codes.Internal, "Failed to save order")
    }

    // Publish order to Pulsar
    if err := PublishOrder(orderModel); err != nil {
        logger.Error("Failed to publish order: ", err)
        return nil, status.Error(codes.Internal, "Failed to publish order")
    }

    logger.Success("Order created successfully: ", orderID)
    return &pb.OrderResponse{
        Status:  "SUCCESS",
        Message: "Order created successfully",
    }, nil
}
