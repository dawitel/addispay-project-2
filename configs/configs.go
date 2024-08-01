package configs

import (
    "os"

    "github.com/dawitel/addispay-project-2/internal/utils"
    
    "gopkg.in/yaml.v2"
)

var logger = utils.GetLogger()



// Config holds the configuration settings for the application.
type Config struct {
    DevelopmentPulsarURL    string `yaml:"dev_pulsar_url"`
    ProductionPulsarURL     string `yaml:"prod_pulsar_url"`
    OrdersTopic             string `yaml:"orders_topic"`
    TransactionsTopic       string `yaml:"transactions_topic"`
    OrdersLogTopic          string `yaml:"orders_log_topic"`
    PaymentsLogTopic        string `yaml:"payments_log_topic"`
    PaymentsLogSubscription string `yaml:"payments-log-subscription"`
    OrdersLogSubscription   string `yaml:"orders-log-subscription"`
    OrderResponseSubscription   string `yaml:"order-response-subscription"`
    GrpcServerAddr          string `yaml:"grpc_server_addr"`
    GRPCPort                string `yaml:"grpc_port"`
    APIGatewayAddr          string `yaml:"api_gateway_addr"`
    FrontendAddr            string `yaml:"frontend_addr"`
    OrderServiceLogFile     string `yaml:"order_service_log_file"`
    PaymentServiceLogFile   string `yaml:"payment_service_log_file"`
    RetryCount              int    `yaml:"retry_count"`
    RetryInterval           int    `yaml:"retry_interval"`
    ExpiryTime              int    `yaml:"expiry_time"` // in seconds
}

// LoadConfig reads the configuration file and unmarshals it into the Config struct.

func LoadConfig() (*Config, error) {
    configPath := "configs/configs.yml";
    // Check if the config file exists
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        return nil, err
    }

    // Read the config file
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, err
    }

    // Unmarshal the config data
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

    // Validate and set default values if necessary
    if config.RetryCount <= 0 {
        config.RetryCount = 3 // Default retry count
    }

    if config.RetryInterval <= 0 {
        config.RetryInterval = 30 // Default retry interval in seconds
    }

    if config.ExpiryTime <= 0 {
        config.ExpiryTime = 120 // Default expiry time in seconds (2 minutes)
    }

    logger.Success("Config loaded successfully from ", configPath)
    return &config, nil
}
