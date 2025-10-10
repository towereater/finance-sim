package queue

import (
	"context"
	"fmt"
	"mainframe-lib/common/config"
	"time"

	"github.com/segmentio/kafka-go"
)

func QueueContent(ctx context.Context, queue config.Queue, topic string, abi string, payload string) error {
	// Create topic writer
	w := &kafka.Writer{
		Addr:                   kafka.TCP(queue.Brokers...),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	// Prepare data for queue
	key := []byte(fmt.Sprintf("%s:%s", abi, time.Now().Format(time.DateTime)))
	value := []byte(payload)

	// Write data on queue
	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)

	return err
}

func UnqueueContent(ctx context.Context, queue config.Queue, topic string) (string, string, error) {
	// Create topic reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: queue.Brokers,
		GroupID: queue.Group,
		Topic:   topic,
	})
	defer r.Close()

	// Read the message from queue
	m, err := r.ReadMessage(ctx)
	if err != nil {
		return "", "", err
	}

	return string(m.Key), string(m.Value), nil
}
