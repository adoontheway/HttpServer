package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var pool *redis.Pool

func InitRedisPool(addr string, password string)  {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c,err := redis.Dial("tcp",addr)
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _,err := c.Do("AUTH",password);err != nil {
					c.Close()
					return nil,err
				}
			}

			if _, err := c.Do("SELECT",0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:3,
		IdleTimeout:time.Minute*10,
	}
}

func Send(cmd string, args ...interface{}) (reply interface{}, err error) {
	conn := pool.Get()
	defer conn.Close()
	reply, err = conn.Do(cmd,args...)
	if err != nil {
		log.Println(err)
	}
	return
}