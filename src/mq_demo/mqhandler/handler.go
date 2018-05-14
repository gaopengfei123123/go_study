package mqhandler

import(
	// "fmt"
	// "mq_demo/mqhandler/kafka"
	"mq_demo/mqhandler/rabbitmq"
)


// MQService mq 服务,目前暂定 kafka, 包含 send 和 read 两个方法
type MQService struct{
	// kafka.Kafka
	rabbitmq.Rabbitmq
}
// MQ 队列应该实现的方法
type MQ interface{
	Read(func(jsonStr []byte))
	Send(key string,value string)
	Delay(key string,value string, expire string)
}