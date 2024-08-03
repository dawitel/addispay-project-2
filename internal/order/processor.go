package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dawitel/addispay-project-2/api/orders/pb"
	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"
)

func ProcessOrderResults(txn *models.Transaction) {
	config, err := configs.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration file: ", err)
	}
	// // donot process the order if the status is "PENDING"
	// if txn.Status == "PENDING" {
	// 	return
	// }
	// Construct the response for the API client
orderResponse := &pb.TransactionResponseForOrder{
    Merchant:      ConvertModelMerchantToPBMerchant(txn.Merchant), // Convert models.Merchant to *pb.Merchant
    TransactionId: txn.TransactionID,
    OrderRequest:  ConvertModelOrderToPBOrderRequest(txn.OrderRequest), // Convert models.Order to *pb.OrderRequest
    Status:        txn.Status,
    Timestamp:     txn.Timestamp,
    Message:       txn.Message,
}
	// write the order response to the DB
	if err = SaveTransactionResponseForOrder(orderResponse); err != nil {
		logger.Error("Failed to write the order response to the Database: ", err)
	}

	// marshal the data to JSON
	orderResponseJSON, err := json.Marshal(orderResponse)
	if err!= nil {
		logger.Error("Failed to marshal the transactions response: ", err)
	}

	URL := fmt.Sprintf("%s/api/v1/processed-transactions", config.APIGatewayAddr)
	
	// send the JSON data to the API gateway
	response, err := http.Post(URL, "application/json", bytes.NewBuffer(orderResponseJSON))
	if err != nil {
		logger.Error("Failed to send the order Response to the API gateway: ", err)
	}

	defer response.Body.Close()
	logger.Success("Order Response data sent to the API gateway: ", response.StatusCode)
}

func ConvertModelMerchantToPBMerchant(modelMerchant models.Merchant) *pb.Merchant {
    products := make([]*pb.Product, len(modelMerchant.Products))
    for i, modelProduct := range modelMerchant.Products {
        products[i] = &pb.Product{
            ProductId:   modelProduct.ProductId,
            ProductName: modelProduct.ProductName,
            ProductPrice:       modelProduct.ProductPrice,
            // Add other fields as needed
        }
    }

    return &pb.Merchant{
        Products:      products,
        MerchantId:    modelMerchant.MerchantId,
        MerchantName:  modelMerchant.MerchantName,
        MerchantEmail: modelMerchant.MerchantEmail,
        MerchantPhone: modelMerchant.MerchantPhone,
        TotalRevenue:  modelMerchant.TotalRevenue,
        TotalProfit:   modelMerchant.TotalProfit,
        NumberOfSales: modelMerchant.NumberOfSales,
    }
}

func ConvertModelOrderToPBOrderRequest(modelOrder models.Order) *pb.OrderRequest {
    return &pb.OrderRequest{
        Merchant:    ConvertModelMerchantToPBMerchant(modelOrder.Merchant),
        CustId:      modelOrder.CustID,
        CustBankAcc: modelOrder.CustBankAcc,
        PhoneNumber: modelOrder.PhoneNumber,
        CustName:    modelOrder.CustName,
        ProductAmount: ConvertModelProductAmountToPB(&modelOrder.ProductAmount),
        TotalAmount:   modelOrder.TotalAmount,
    }
}

func ConvertModelProductAmountToPB(modelProductAmount *models.ProductAmount) *pb.ProductAmount {
    return &pb.ProductAmount{
        Product: ConvertModelProductToPBProduct(&modelProductAmount.Product),
        ProductAmount: modelProductAmount.ProductAmount,
    }
}

func ConvertModelProductToPBProduct(modelProduct *models.Product) *pb.Product {
    return &pb.Product{
        ProductId: modelProduct.ProductId,
        ProductName: modelProduct.ProductName,
        ProductPrice: modelProduct.ProductPrice,
    }
}



