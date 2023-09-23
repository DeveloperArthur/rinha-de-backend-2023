package caching

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golang-first-api-rest/models"
	"golang-first-api-rest/util"
	"time"
)

func ClientSingleton() (*redis.Client, context.Context) {
	var instancia *redis.Client
	if instancia == nil {
		instancia = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return instancia, context.Background()
}

func Set(pessoa *models.Pessoa, key string) {
	instancia, ctx := ClientSingleton()
	expiration := 10 * time.Minute
	pessoaEmJson := util.Serialize(pessoa)
	err := instancia.Set(ctx, key, pessoaEmJson, expiration).Err()
	if err != nil {
		panic(err)
	}
}

func SetList(pessoas *[]models.Pessoa, key string) {
	instancia, ctx := ClientSingleton()
	expiration := 3 * time.Minute
	pessoaEmJson := util.SerializeList(pessoas)
	err := instancia.Set(ctx, key, pessoaEmJson, expiration).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	instancia, ctx := ClientSingleton()
	return instancia.Get(ctx, key).Result()
}
