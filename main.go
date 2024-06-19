package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/navikt/go-leesah"
)

func main() {
	os.Setenv("KAFKA_GROUP_ID", uuid.New().String())

	var config leesah.RapidConfig
	flag.StringVar(&config.Brokers, "brokers", os.Getenv("KAFKA_BROKERS"), "Kafka broker")
	flag.StringVar(&config.Topic, "topic", os.Getenv("KAFKA_TOPIC"), "Kafka topic")
	flag.StringVar(&config.GroupID, "group-id", os.Getenv("KAFKA_GROUP_ID"), "Kafka group ID")
	flag.StringVar(&config.KafkaCertPath, "kafka-cert-path", os.Getenv("KAFKA_CERTIFICATE_PATH"), "Path to Kafka certificate")
	flag.StringVar(&config.KafkaPrivateKeyPath, "kafka-private-key-path", os.Getenv("KAFKA_PRIVATE_KEY_PATH"), "Path to Kafka private key")
	flag.StringVar(&config.KafkaCAPath, "kafka-ca-path", os.Getenv("KAFKA_CA_PATH"), "Path to Kafka CA certificate")
	config.Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	rapid, err := leesah.NewRapid("team-teamwork", config)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create rapid: %s", err))
		return
	}
	defer rapid.Close()

	slog.Info("rapid is running")
	if err := rapid.Run(Answer); err != nil {
		slog.Error(fmt.Sprintf("failed to run rapid: %s", err))
	}
}

func Answer(question leesah.Question, log *slog.Logger) (string, bool) {
	log.Info(fmt.Sprintf("%+v", question))

	switch question.Category {
	case "team-registration":
		return "#371F76", true
	}

	return "", false
}
