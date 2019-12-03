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
	client.Set("count", 1, 0)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go reduce4(client, &wg)
	}
	wg.Wait()

	fmt.Println("done")
}

func reduce4(client *redis.Client, wg *sync.WaitGroup) {
	ok, _ := lock(client)
	if ok {
		count, _ := client.Get("count").Int()
		if count > 0 {
			client.Decr("count")
		}
		unlock(client)
	}
	wg.Done()
}

func lock(client *redis.Client) (bool, error) {
	return client.SetNX("count.lock", true, 0).Result()
}

func unlock(client *redis.Client) {
	client.Del("count.lock")
}
