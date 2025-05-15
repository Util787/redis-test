package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name string
	Age  int
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	user := User{Name: "Alice", Age: 30}

	err := rdb.HSet(ctx, "user:1", "name", user.Name, "age", user.Age).Err()
	if err != nil {
		log.Fatalf("Error setting value in Redis: %v", err)
	}

	user2 := User{Name: "Alice2", Age: 32}
	userJSON, err := json.Marshal(user2)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	err = rdb.Set(ctx, "user:2", userJSON, 0).Err()
	if err != nil {
		log.Fatalf("Error setting value in Redis: %v", err)
	}

	// Получение данных из HSET
	name, err := rdb.HGet(ctx, "user:1", "name").Result()
	if err != nil {
		log.Fatalf("Error getting name from Redis: %v", err)
	}

	ageStr, err := rdb.HGet(ctx, "user:1", "age").Result()
	if err != nil {
		log.Fatalf("Error getting age from Redis: %v", err)
	}

	fmt.Println("user:1 HGET>", name, ageStr)

	userFields, err := rdb.HGetAll(ctx, "user:1").Result()
	if err != nil {
		log.Fatalf("Error getting all fields from Redis: %v", err)
	}
	fmt.Println()
	fmt.Println("user:1 HGETALL> ", userFields)
	fmt.Println()

	for k, _ := range userFields {

		fmt.Println("key from userFields:", k)
	}

	// Получение данных из строки
	userJSONStr, err := rdb.Get(ctx, "user:2").Result()
	if err != nil {
		log.Fatalf("Error getting user JSON from Redis: %v", err)
	}

	fmt.Println()
	fmt.Println("user:2 >", userJSONStr)

}
