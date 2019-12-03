package main

import (
	"fmt"
	"sync"
	"time"

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
	lock(client, func() {
		count, _ := client.Get("count").Int()
		if count > 0 {
			client.Decr("count")
		}
	})
	wg.Done()
}

func lock(client *redis.Client, doSomething func()) {
	ok, _ := client.SetNX("count.lock", true, 60*time.Second).Result()
	if ok {
		doSomething()
		unlock(client)
	}
}

func unlock(client *redis.Client) {
	client.Del("count.lock")
}
