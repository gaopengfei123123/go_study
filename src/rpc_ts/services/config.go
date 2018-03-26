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
