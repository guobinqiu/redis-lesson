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
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go reduce4(client, &wg, &m)
	}
	wg.Wait()

	fmt.Println("done")
}

func reduce4(client *redis.Client, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	count, _ := client.Get("count").Int()
	if count > 0 {
		client.Decr("count")
	}
	m.Unlock()
	wg.Done()
}
