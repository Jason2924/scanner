package databases

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Jason2924/scanner/backend/config"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

type IRedisCache interface {
	Connect() *redis.Client
	Close() error
	Ping(ctxt context.Context) error
	Store(ctxt context.Context, name string, item interface{}, expr time.Duration) error
	Retrieve(ctxt context.Context, name string, item interface{}) (bool, error)
}

type redisCache struct {
	address  string
	username string
	password string
	database int
}

func NewRedisCache(options *config.ConfigRedis) IRedisCache {
	return &redisCache{
		address:  options.Address,
		username: options.Username,
		password: options.Password,
		database: 0,
	}
}

func (cac *redisCache) Connect() *redis.Client {
	redisOnce.Do(func() {
		options := redis.Options{
			Addr:     cac.address,
			Password: cac.password,
			DB:       cac.database,
		}
		redisClient = redis.NewClient(&options)
	})
	return redisClient
}

func (cac *redisCache) Close() error {
	if redisClient == nil {
		return nil
	}
	return redisClient.Close()
}

func (cac *redisCache) Ping(ctxt context.Context) error {
	_, erro := cac.Connect().Ping(ctxt).Result()
	return erro
}

func (cac *redisCache) Store(ctxt context.Context, name string, item interface{}, expr time.Duration) error {
	client := cac.Connect()
	byteItem, erro := json.Marshal(item)
	if erro != nil {
		return fmt.Errorf("Error occured while marshalling item: %v", erro.Error())
	}
	rsul := client.Set(ctxt, name, byteItem, expr)
	if rsul.Err() != nil {
		return fmt.Errorf("Error occured while storing to cache: %v", rsul.Err())
	}
	return nil
}

func (cac *redisCache) Retrieve(ctxt context.Context, name string, item interface{}) (bool, error) {
	client := cac.Connect()
	result := []byte{}
	erro := client.Get(ctxt, name).Scan(&result)
	if erro == redis.Nil {
		return false, nil
	} else if erro != nil {
		return false, fmt.Errorf("Error occured while retrieving to cache: %v", erro.Error())
	}
	erro = json.Unmarshal(result, item)
	if erro != nil {
		return false, fmt.Errorf("Error occured while unmarshalling item: %v", erro.Error())
	}
	return true, nil
}
