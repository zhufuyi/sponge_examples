package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/zhufuyi/sponge/pkg/kafka"
	"github.com/zhufuyi/sponge/pkg/logger"

	orderV1 "order/api/order/v1"
	"order/internal/config"
	"order/internal/service"
)

var (
	cg   *kafka.ConsumerGroup
	once sync.Once
)

func InitConsumerGroup() error {
	var err error
	cg, err = kafka.InitConsumerGroup(
		config.Get().Kafka.Brokers,
		config.Get().Kafka.GroupID,
		kafka.ConsumerWithVersion(sarama.V3_6_0_0),
		kafka.ConsumerWithClientID("eshop-order"),
		kafka.ConsumerWithZapLogger(logger.Get()),
	)
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(time.Second * 5)
		for {
			err := cg.Consume(context.Background(), []string{config.Get().Kafka.OrderTopic}, submitOrder())
			logger.Infof("kafka consumer group exited, err: %v", err)
		}
	}()

	return err
}

// GetConsumerGroup get kafka consumer group
func GetConsumerGroup() *kafka.ConsumerGroup {
	if cg == nil {
		once.Do(func() {
			err := InitConsumerGroup()
			if err != nil {
				panic(fmt.Errorf("initialize kafaka consumer group failed: %v", err))
			}
		})
	}

	return cg
}

// CloseConsumerGroup close kafka consumer group
func CloseConsumerGroup() error {
	if cg == nil {
		return nil
	}

	return cg.Close()
}

// submit order message to order server
func submitOrder() func(msg *sarama.ConsumerMessage) error {
	s := service.NewOrderServer()

	return func(msg *sarama.ConsumerMessage) error {
		defer func() {
			if e := recover(); e != nil {
				logger.Error("submit order panic", logger.Any("err", e))
			}
		}()

		logger.Info("received order message from kafka",
			logger.String("topic", msg.Topic),
			logger.Int("partition", int(msg.Partition)),
			logger.Int64("offset", msg.Offset),
			logger.String("key", string(msg.Key)),
			logger.String("value", string(msg.Value)),
		)
		req := &orderV1.SubmitOrderRequest{}
		err := json.Unmarshal(msg.Value, req)
		if err != nil {
			return err
		}

		logger.Info("debug req", logger.Any("req", req))

		ctx := context.Background()
		_, err = s.Submit(ctx, req)

		return err
	}
}

// only show message
func createOrderHandler(msg *sarama.ConsumerMessage) error {
	logger.Info("received order message from kafka",
		logger.String("topic", msg.Topic),
		logger.Int("partition", int(msg.Partition)),
		logger.Int64("offset", msg.Offset),
		logger.String("key", string(msg.Key)),
		logger.String("value", string(msg.Value)),
	)
	req := &orderV1.SubmitOrderRequest{}
	err := json.Unmarshal(msg.Value, req)
	if err != nil {
		return err
	}

	logger.Debug("debug req", logger.Any("req", req))

	return nil
}
