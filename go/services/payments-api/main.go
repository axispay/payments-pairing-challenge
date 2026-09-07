package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"payments-api/api/rest"
	"payments-api/external/db"
	externalkafka "payments-api/external/kafka"
	"payments-api/internal/account"
	"payments-api/internal/payment"
	"payments-api/internal/transfer"
)

func main() {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	mongoURI := getenv("MONGODB_URI", "mongodb://localhost:27017/payments")
	mongo, err := db.Connect(mongoURI)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to MongoDB")
	}
	defer mongo.Disconnect()

	producer := externalkafka.NewProducer(os.Getenv("KAFKA_BROKERS"))
	defer producer.Close()

	accountRepo := account.NewRepository(mongo.DB())
	paymentRepo := payment.NewRepository(mongo.DB())
	transferRepo := transfer.NewRepository(mongo.DB())

	accountService := account.NewService(accountRepo)
	paymentService := payment.NewService(paymentRepo, producer)
	transferService := transfer.NewService(transferRepo, producer)

	h := rest.New(accountService, paymentService, transferService)
	router := gin.Default()
	rest.RegisterRoutes(router, h)

	port := getenv("PORT", "3001")
	log.Info().Str("port", port).Msg("payments-api listening")
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Err(err).Msg("server error")
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
