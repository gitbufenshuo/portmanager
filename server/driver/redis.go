package driver

import (
	"github.com/go-redis/redis"
)

type RedisConfig struct {
	HostPort string
	Password string
	_session *redis.Client
}

func (rc *RedisConfig) redis_init() error {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.HostPort,
		Password: rc.Password, // no password set
		DB:       1,           // 我们只使用 1
	})

	if _, err := client.Ping().Result(); err != nil {
		return err
	}
	rc._session = client
	return nil
}

func (rc *RedisConfig) GetRedisSession() *redis.Client {
	return rc._session
}

func Init(rc *RedisConfig) error {
	return rc.redis_init()
}
