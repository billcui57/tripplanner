package main

import (
	tripplancontroller "github/billcui57/tripplanner/Controllers/TripplanController"
	utils "github/billcui57/tripplanner/Utils"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v8"
	libredis "github.com/go-redis/redis/v8"

	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func main() {

	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted("10-H")
	if err != nil {
		log.Fatal(err)
		return
	}
	// Create a redis client.
	client := libredis.NewClient(&redis.Options{
		Addr:     utils.GetEnvVarOrDefault("REDIS_ADDR", "redis:6379"),
		Password: utils.GetEnvVarOrDefault("REDIS_PASSWORD", ""),
		DB:       0,
	})

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new middleware with the limiter instance.
	rateLimitMiddleware := mgin.NewMiddleware(limiter.New(store, rate))

	engine := gin.Default()
	engine.ForwardedByClientIP = true

	engine.Use(cors.Default())
	engine.Use(rateLimitMiddleware)

	engine.POST("/plan-trip", tripplancontroller.Plantrip)

	engine.Run()
}
