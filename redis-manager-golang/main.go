package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, key := range keys {
		fmt.Printf("key: %s\n", key)
		fmt.Printf("type: %T\n", key)

		val, err := redis.String(conn.Do("GET", key))
		if err != nil {
			panic(err)
		} else {
			fmt.Println("key", val)
		}
	}
}
