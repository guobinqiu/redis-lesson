package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "lucy654123xlhb",
		DB:       0,
	})
	client.Set("count", 1000, 0)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go reduce3(client, &wg)
	}
	wg.Wait()

	fmt.Println("done")
}

func reduce3(client *redis.Client, wg *sync.WaitGroup) {
	client.Decr("count")
	wg.Done()
}
