package main

import (
	"context"
	"log"
	"os"
	"praisindo/entity"
	"time"

	"golang.org/x/exp/slog"

	"github.com/IBM/sarama"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		// Process the message as per your requirement here
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	brokers := []string{entity.KafkaBrokerAddress}
	groupID := entity.DefaultGroupID
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0 // specify appropriate Kafka version
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		slog.Error("error when setup sarama.NewConsumerGroup:", slog.Any("error", err))
		os.Exit(1)
	}

	ctx := context.Background()
	slog.InfoContext(ctx, "Start consuming from topic", slog.Any("topic", entity.DefaultTopic))
	for {
		if err := consumerGroup.Consume(ctx, []string{entity.DefaultTopic}, exampleConsumerGroupHandler{}); err != nil {
			slog.Error("error when call consumerGroup.Consume:", slog.Any("error", err))
		}
	}

}
