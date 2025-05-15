package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	ID         string
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Age        int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("failed to ping", err)
		return
	}

	fmt.Println(ping)

	ElliotId := uuid.NewString()
	jsonStr, err := json.Marshal(Person{
		ID:         ElliotId,
		Name:       "Elliot",
		Age:        24,
		Occupation: "QA",
	})
	if err != nil {
		fmt.Println("failed marshal")
		return
	}

	ElliotKey := fmt.Sprintf("person:%s", ElliotId)

	err = client.Set(context.Background(), ElliotKey, jsonStr, 0).Err()
	if err != nil {
		fmt.Println("failed to set value", err)
		return
	}

	val0, err := client.Get(context.Background(), ElliotKey).Result()
	if err != nil {
		fmt.Println("failed to get person from redis")
		return
	}
	fmt.Printf("type:%T\n Result:%s\n", val0, val0)

	var testP Person
	err = json.Unmarshal([]byte(val0), &testP)
	if err != nil {
		fmt.Println("failed unmarshal")
		return
	}
	fmt.Printf("type:%T\n Result:%v\n", testP, testP)

}
