package sender

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/zhufuyi/sponge/pkg/kafka"

	"order/internal/config"
)

var (
	syncProducer *kafka.SyncProducer
	once         sync.Once
)

// InitSyncProducer initializes the Kafka producer
func InitSyncProducer() error {
	var err error
	syncProducer, err = kafka.InitSyncProducer(config.Get().Kafka.Brokers, kafka.SyncProducerWithVersion(sarama.V3_6_0_0))
	return err
}

// GetSyncProducer get sync producer client
func GetSyncProducer() *kafka.SyncProducer {
	if syncProducer == nil {
		once.Do(func() {
			err := InitSyncProducer()
			if err != nil {
				panic(fmt.Errorf("init kafaka producer failed: %v", err))
			}
		})
	}

	return syncProducer
}

// CloseSyncProducer close the Kafka producer
func CloseSyncProducer() error {
	if syncProducer == nil {
		return nil
	}
	return syncProducer.Close()
}
