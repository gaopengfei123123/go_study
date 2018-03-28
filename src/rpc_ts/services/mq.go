package services

import(
	"github.com/segmentio/kafka-go"
	"fmt"
	"context"
	// "time"
	"github.com/astaxie/beego/logs"
)

func init(){
	fmt.Println("初始化 log 配置")
	// log 开异步
	logs.Async(1e3)
	config := fmt.Sprintf(`{"filename":"%s","separate":["error", "warning", "notice", "info", "debug"]}`, LOG_PATH )
	logs.SetLogger(logs.AdapterMultiFile, config)
}

// MQService mq 服务,目前暂定 kafka, 包含 send 和 read 两个方法
type MQService struct{}

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

	logs.Debug("already send msg, key:", key)
}

func (mq *MQService) Read(){
	logs.Debug("mq reader is starting")
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
		logs.Info("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		ServerService(m.Value)

		r.CommitMessages(ctx, m)
	}
	defer r.Close()
}