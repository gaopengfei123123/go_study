package services

import(
	"github.com/segmentio/kafka-go"
	"fmt"
	"context"
	// "time"
)

// MQService mq 服务,目前暂定 kafka, 包含 send 和 read 两个方法
type MQService struct{}

const (
	// KAFKA_TOPIC 指定的主题名
	KAFKA_TOPIC = "my-topic"
	// KAFKA_PARTITION 指定分区
	KAFKA_PARTITION = 0
	// KAFKA_HOST kafka 链接地址
	KAFKA_HOST = "localhost:9092"
)

// Send 向队列插入数据 (目前是使用 kafka)
func (mq *MQService) Send(key string,value string){
	conn, _ := kafka.DialLeader(context.Background(), "tcp", KAFKA_HOST, KAFKA_TOPIC, KAFKA_PARTITION)
	defer conn.Close()

	conn.WriteMessages(
		kafka.Message{
			Key: []byte(key),
			Value: []byte(value),
		},
	)

	fmt.Println("already send msg")
}

func (mq *MQService) Read(){
	fmt.Println("mq reader is starting")
	startReading()
}

// 开始监听消息队列
func startReading(){
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{KAFKA_HOST},
		Topic:     KAFKA_TOPIC,
		Partition: KAFKA_PARTITION,
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