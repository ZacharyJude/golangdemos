package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if "" != password {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool          *redis.Pool
	redisServer   = flag.String("redisServer", "localhost:6379", "")
	redisPassword = flag.String("redisPassword", "", "")
)

func main() {
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword)
	defer pool.Close()

	key := "demo---"
	c := pool.Get()
	defer c.Close()

	//reply, err := redis.Strings(c.Do("keys", fmt.Sprintf("%s*", key)))
	args := redis.Args{}.Add(key).Add(10).Add(23).Add(34).Add(45)
	reply, err := redis.Int64(c.Do("sadd", args...))
	if nil != err {
		fmt.Printf("do error %s\n", err)
		return
	}

	fmt.Printf("reply %d\n", reply)

	args = redis.Args{}.Add(key)
	smembers, err := redis.Values(c.Do("smembers", args...))
	if nil != err {
		fmt.Printf("do error %s\n", err)
		return
	}

	for _, e := range smembers {
		fmt.Printf("%s ", e)
	}
	fmt.Printf("\n")
}
