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
		reduce2(client, &wg)
	}
	wg.Wait()

	fmt.Println("done")
}

//https://redis.io/topics/transactions
func reduce2(client *redis.Client, wg *sync.WaitGroup) {
	pipe := client.TxPipeline()
	i, _ := client.Get("count").Int()
	pipe.Set("count", i-1, 0)
	pipe.Exec()

	wg.Done()
}
