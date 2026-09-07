package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"merchant-analytics/api/rest"
	"merchant-analytics/external/db"
	externalkafka "merchant-analytics/external/kafka"
	"merchant-analytics/internal/analytics"
	"os"
)

func main() {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	mongoURI := getenv("MONGODB_URI", "mongodb://localhost:27017/payments")
	mongo, err := db.Connect(mongoURI)
	if err != nil {
		log.Fatal().Err(err).Msg("mongo connect failed")
	}
	defer mongo.Disconnect()
	repo := analytics.NewRepository(mongo.DB())
	svc := analytics.NewService(repo)
	h := rest.New(svc)
	go externalkafka.NewConsumer(os.Getenv("KAFKA_BROKERS"), mongo.DB()).Start()
	r := gin.Default()
	rest.RegisterRoutes(r, h)
	port := getenv("PORT", "3002")
	log.Info().Str("port", port).Msg("merchant-analytics listening")
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Err(err).Msg("server error")
	}
}
func getenv(k, f string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return f
}
