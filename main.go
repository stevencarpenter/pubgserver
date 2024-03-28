package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	// Initialize the Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-c82bfc2f:6379",
		Password: "password",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong, "Connected to Redis")
}

func getAccountData(w http.ResponseWriter, r *http.Request) {
	accountId := r.URL.Query().Get("accountId")
	if accountId == "" {
		http.Error(w, "accountId is required", http.StatusBadRequest)
		return
	}

	data, err := redisClient.Get(ctx, accountId).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		log.Printf("Error fetching data from Redis: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{accountId: data}); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/pubg/leaderboard", getAccountData)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
