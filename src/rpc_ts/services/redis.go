package services

import(
	"github.com/go-redis/redis"
)

// RedisClient 服务主体
type RedisClient struct{
	client *redis.Client
}

var redisClient *RedisClient

// NewRedisClient 生成一个 redis 客户端
func NewRedisClient() *RedisClient{
	if redisClient == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		redisClient.client = client
	}

	return redisClient
}


// Get 通过 key 获取缓存
func (rd *RedisClient) Get(key string) string{
	val, err := rd.client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// Set 设置缓存
func (rd *RedisClient) Set(key string,value string) bool{
	err := rd.client.Set(key,value,REDIS_CACHE_TTL)
	if err != nil {
		panic(err)
	}
	return true
}