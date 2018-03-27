package services

// redis 的一些配置
const (
	REDIS_HOST = "localhost:6379"
	REDIS_PASSWORD = ""
	REDIS_DB = 0
	REDIS_CACHE_TTL = 7200
)


// kafka 的一些配置
const (
	// KAFKA_TOPIC 指定的主题名
	KAFKA_TOPIC = "my-topic"
	// KAFKA_PARTITION 指定分区
	KAFKA_PARTITION = 0
	// KAFKA_HOST kafka 链接地址
	KAFKA_HOST = "localhost:9092"
)

// 消费者服务的一些配置
const (
	// MAX_EXEC_NUM 最大执行次数
	MAX_EXEC_NUM = 3
	// MAX_EXEC_TIME 最大执行时间(单位 秒)
	MAX_EXEC_TIME = 20
)
