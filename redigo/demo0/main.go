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

	//key := "lbsplay:houses:"
	key := "lbsplay:houses:364"
	//keys := []string{"lbsplay:houses:poiUID:2ca237afb2b1b6bf31870a9a"}
	c := pool.Get()
	defer c.Close()

	//reply, err := redis.Strings(c.Do("keys", fmt.Sprintf("%s*", key)))
	reply, err := redis.StringMap(c.Do("hgetall", key))
	/*
		reply, err := redis.Strings(
			c.Do(
				"mget",
				redis.Args{}.AddFlat(keys)...))
	*/
	if nil != err {
		fmt.Println(err)
		return
	}

	for k, v := range reply {
		/*
			if "" == s {
				fmt.Println("gotcha!")
			}
		*/
		fmt.Printf("%s\t%s\n", k, v)
	}
}
