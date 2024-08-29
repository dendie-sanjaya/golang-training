package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/signal"
	"praisindo_consumer_1/config"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

const (
	basedir  = "."
	fileName = "data.csv"
)

func main() {
	// Konfigurasi konsumer grup Kafka
	configKakfa := sarama.NewConfig()
	configKakfa.Consumer.Return.Errors = true
	configKakfa.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	configKakfa.Net.DialTimeout = 10 * time.Second
	configKakfa.Net.ReadTimeout = 10 * time.Second
	configKakfa.Net.WriteTimeout = 10 * time.Second

	// Inisialisasi konsumer grup
	brokers := []string{config.KafkaBroker}
	topic := config.KafkaTopicTransfer
	groupID := config.KafkaGroupID
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, configKakfa)
	if err != nil {
		log.Fatalf("Failed to start consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// Buat channel untuk menangani sinyal sistem
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Buat handler untuk konsumer grup
	handler := &ConsumerGroupHandler{}

	// Konsumsi pesan dari topik
	go func() {
		for {
			if err := consumerGroup.Consume(context.Background(), []string{topic}, handler); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
		}
	}()

	// Tunggu sinyal sistem untuk keluar
	<-signals
	log.Println("Consumer stopped")
}

// ConsumerGroupHandler adalah handler untuk konsumer grup
type ConsumerGroupHandler struct {
	csvWriter *csv.Writer
}

// Setup dijalankan saat konsumer grup diinisialisasi
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	// Buka atau buat file CSV
	file, err := os.OpenFile("log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Inisialisasi CSV writer
	h.csvWriter = csv.NewWriter(file)

	// Tulis header CSV jika file baru dibuat
	/*
		if fileInfo, err := file.Stat(); err == nil && fileInfo.Size() == 0 {
			header := []string{}
			if err := h.csvWriter.Write(header); err != nil {
				return err
			}
			h.csvWriter.Flush()
		}
	*/

	return nil
}

// Cleanup dijalankan saat konsumer grup dihentikan
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	if h.csvWriter != nil {
		h.csvWriter.Flush()
	}
	return nil
}

// ConsumeClaim menangani pesan yang dikonsumsi
func (h *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Consumed message: %s\n", string(msg.Value))
		sess.MarkMessage(msg, "")

		// Tulis pesan ke file CSV
		record := []string{string(msg.Value)}
		if err := h.csvWriter.Write(record); err != nil {
			log.Printf("Error writing to CSV: %v", err)
		}
		h.csvWriter.Flush()

		sess.MarkMessage(msg, "")
	}
	return nil
}
