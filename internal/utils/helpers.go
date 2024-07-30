package utils

import (
    "fmt"
    "os"
    "math/rand"
    "time"

    
    "github.com/joho/godotenv"
)

// GenerateID generates a random ID for transactions.
func GenerateID() string {
    return fmt.Sprintf("%d", time.Now().UnixNano())
}

// RandomSuccess simulates a random success/failure scenario.
func RandomSuccess() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(10) < 7 // 70% success rate
}

func GoDotEnvVariable(key string) string {
  err := godotenv.Load(".env")
  if err != nil {
    logger.Error("Error loading .env file")
  }

  return os.Getenv(key)
}