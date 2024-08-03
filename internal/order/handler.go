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
    // Basic Validation for the request
    if req.CustId == "" || req.TotalAmount <= 0 {
        return nil, status.Error(codes.InvalidArgument, "Invalid order details")
    }

    // Convert gRPC request to internal model
    orderModel := &models.Order{
        Merchant:    ConvertPBMerchantToModelMerchant(req.Merchant), // TODO: -> type conversion
        OrderID:     orderID,
        CustID:      req.CustId,
        CustBankAcc: req.CustBankAcc,
        PhoneNumber: req.PhoneNumber,
        CustName:    req.CustName,
        ProductAmount: ConvertPBProductAmountToProduct(req.ProductAmount), // TODO -> type conversion
        TotalAmount: req.TotalAmount,
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

// ConvertPBMerchantToModelMerchant converts the protocol buffer merchant to the internal merchant model.
func ConvertPBMerchantToModelMerchant(pbMerchant *pb.Merchant) models.Merchant {
    // type coversion for the products without conflict
    products := make([]models.Product, len(pbMerchant.Products))
    for i, pbProduct := range pbMerchant.Products {
        products[i] = ConvertPBProductToModelProduct(pbProduct)
    }
    return models.Merchant{
        Products: products,
        MerchantId: pbMerchant.MerchantId,
        MerchantName: pbMerchant.MerchantName,
        MerchantEmail: pbMerchant.MerchantEmail,
        MerchantPhone: pbMerchant.MerchantPhone,
        TotalRevenue: pbMerchant.TotalRevenue,
        TotalProfit: pbMerchant.TotalProfit,
        NumberOfSales: pbMerchant.NumberOfSales,
    }
}

func ConvertPBProductToModelProduct(pbProduct *pb.Product) models.Product {
    return models.Product{
        ProductId: pbProduct.ProductId,
        ProductName: pbProduct.ProductName,
        ProductPrice: pbProduct.ProductPrice,
    }
}

func ConvertPBProductAmountToProduct(pbProductAmount *pb.ProductAmount) models.ProductAmount {
    return models.ProductAmount{
        Product: ConvertPBProductToModelProduct(pbProductAmount.Product),
        ProductAmount: pbProductAmount.ProductAmount,
    }
}