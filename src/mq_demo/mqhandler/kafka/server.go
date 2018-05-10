package kafka

import(
	"github.com/segmentio/kafka-go"
	"context"
	"github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/logs"
	// "fmt"
)

// kafka 的一些配置
const (
	// KafkaTopic 指定的主题名
	KafkaTopic = "test-topic"
	// KafkaPatition 指定分区
	KafkaPatition = 0
	// KafkaHost kafka 链接地址
	KafkaHost = "localhost:9092"
)

// Kafka 的实体
type Kafka struct{}

// Send 向队列发送消息
func (kf Kafka) Send(key string,value string){
	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaHost, KafkaTopic, KafkaPatition)
	checkErr(err)
	defer conn.Close()

	conn.WriteMessages(
		kafka.Message{
			Key: []byte(key),
			Value: []byte(value),
		},
	)
	logs.Debug("already send msg, key:", key)
}

// Read 读取队列信息
func (kf Kafka) Read(f func(jsonStr []byte)){
	logs.Debug("mq reader is starting")
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{KafkaHost},
		Topic:     KafkaTopic,
		Partition: KafkaPatition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		logs.Info("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		go f(m.Value)

		r.CommitMessages(ctx, m)
	}
	defer r.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}