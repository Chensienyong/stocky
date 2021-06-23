package main

import (
	"fmt"
	"github.com/chensienyong/stocky"
	"github.com/chensienyong/stocky/connection"
	"github.com/chensienyong/stocky/database"
	"github.com/chensienyong/stocky/handler"
	"github.com/chensienyong/stocky/middleware"
	"github.com/chensienyong/stocky/pkg/redis"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	dbOpt := database.Option{
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		Host:         os.Getenv("POSTGRES_HOST"),
		Port:         os.Getenv("POSTGRES_PORT"),
		Database:     os.Getenv("POSTGRES_DATABASE"),
		Timezone:     os.Getenv("POSTGRES_TIMEZONE"),
	}

	postgres, _ := database.NewPostgres(dbOpt)

	redisOpt := redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}

	redis, _ := redis.NewRedis(redisOpt)

	alphaVantage := connection.NewAlphaVantage(os.Getenv("ALPHAVANTAGE_HOST"), os.Getenv("ALPHAVANTAGE_API_KEY"), 3*time.Second)

	stck := stocky.NewStocky(postgres, redis, alphaVantage)

	handler := handler.NewHandler(stck)

	router := httprouter.New()

	decorator := []middleware.Decorator{middleware.WithBasicAuth(os.Getenv("STOCKY_USER"), os.Getenv("STOCKY_PASSWORD"))}
	router.GET("/healthz", handler.Healthz)
	router.GET("/stocks", apply(handler.GetStocks, decorator...))
	router.GET("/time-series/:stock/dailies", apply(handler.FetchDailyTimeSeries, decorator...))

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	log.Println("Listening at port", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), co.Handler(router)))
}

func apply(f middleware.HandleWithError, decorators ...middleware.Decorator) httprouter.Handle {
	return middleware.HTTP(middleware.ApplyDecorators(f, decorators...))
}
