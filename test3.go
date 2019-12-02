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
	client.Set("count", 0, 0)

	var wg sync.WaitGroup
	no := 1000
	for i := 0; i < no; i++ {
		wg.Add(1)
		go run3(client, &wg)
	}
	wg.Wait()

	fmt.Println("done")
}

func run3(client *redis.Client, wg *sync.WaitGroup) {
	client.Incr("count")
	wg.Done()
}
