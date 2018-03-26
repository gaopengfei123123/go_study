package main

import(
	"github.com/segmentio/kafka-go"
	"context"
	"fmt"
	// "time"
)

func main() {
	// to consume messages
	topic := "my-topic"
	partition := 0


	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		r.CommitMessages(ctx, m)
	}
	fmt.Println("done")

	defer r.Close()
}