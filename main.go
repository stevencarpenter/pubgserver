package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
)

var ctx = context.Background()
var rdb *redis.Client

var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

func init() {
	// Initialize the Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-c82bfc2f:6379",
		Password: REDIS_PASSWORD,
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
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

	data, err := rdb.HGetAll(ctx, accountId).Result()
	if errors.Is(err, redis.Nil) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		log.Printf("Error fetching data from Redis: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/pubg/leaderboard", getAccountData)

	fmt.Println("Server is running on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
