package driver

import (
	"github.com/go-redis/redis"
)

type RedisConfig struct {
	HostPort string
	Password string
	Database int
	_session *redis.Client
}

func (rc *RedisConfig) redis_init() error {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.HostPort,
		Password: rc.Password,
		DB:       rc.Database,
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
func (rc *RedisConfig) Close() {
	rc._session.Close()
}

var RedisConf RedisConfig

func Init(rc *RedisConfig) error {
	return rc.redis_init()
}
