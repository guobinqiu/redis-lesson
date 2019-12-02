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
	client.Set("count", 0, 0).Err()

	var wg sync.WaitGroup
	no := 1000
	for i := 0; i < no; i++ {
		wg.Add(1)
		run2(client, &wg)
	}

	wg.Wait()

	fmt.Println("done")
}

//https://redis.io/topics/transactions
func run2(client *redis.Client, wg *sync.WaitGroup) {
	pipe := client.TxPipeline()
	i, _ := client.Get("count").Int()
	pipe.Set("count", i+1, 0)
	pipe.Exec()

	wg.Done()
}
