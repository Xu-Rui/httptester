package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// TODO 之后可以添加连接池

type redisClient struct {
	conn redis.Conn
	pool redis.Pool
}

var RedisClient redisClient

func InitCodis(url string) error {
	conn, err := redis.Dial("tcp", url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	RedisClient.conn = conn
	return err
}

func (r *redisClient) Get(key string) (string, error) {

	result, err := redis.String(r.conn.Do("get", key))
	if err != nil {
		fmt.Println(err)
	}

	return result, nil
}

func (r *redisClient) Set(key, value string) error {

	_, err := r.conn.Do("set", key, value)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
