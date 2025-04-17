package databases

import (
	"context"
	"testing"
	"time"

	"github.com/Jason2924/scanner/backend/config"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Log("Hello")
	redisConfig := config.ConfigRedis{
		Address:  "localhost:6379",
		Username: "",
		Password: "",
	}
	redisCache := NewRedisCache(&redisConfig)
	ctxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// test string
	reqtStr := "hello"
	erro := redisCache.Store(ctxt, "cache:test-string", reqtStr, 0)
	require.NoError(t, erro)
	respStr := ""
	_, erro = redisCache.Retrieve(ctxt, "cache:test-string", &respStr)
	require.NoError(t, erro)
	require.Equal(t, respStr, reqtStr)
	respStrErr := ""
	_, erro = redisCache.Retrieve(ctxt, "cache:test-string-error", &respStrErr)
	require.Nil(t, erro)
	// test slice
	reqtSli := []string{
		"hello",
		"world",
	}
	erro = redisCache.Store(ctxt, "cache:test-slice", reqtSli, 0)
	require.NoError(t, erro)
	respSli := []string{}
	_, erro = redisCache.Retrieve(ctxt, "cache:test-slice", &respSli)
	require.NoError(t, erro)
	require.EqualValues(t, reqtSli, respSli)
	respSliErr := []string{}
	_, erro = redisCache.Retrieve(ctxt, "cache:test-slice-error", &respSliErr)
	require.Nil(t, erro)
	// test map
	reqtMap := map[string]bool{
		"hello": true,
		"world": false,
	}
	erro = redisCache.Store(ctxt, "cache:test-map", reqtMap, 0)
	require.NoError(t, erro)
	respMap := make(map[string]bool)
	_, erro = redisCache.Retrieve(ctxt, "cache:test-map", &respMap)
	require.NoError(t, erro)
	require.EqualValues(t, reqtMap, respMap)
	respMapErr := make(map[string]bool)
	_, erro = redisCache.Retrieve(ctxt, "cache:test-map-error", &respMapErr)
	require.Nil(t, erro)
	// struct
	type reqtStruct struct {
		FieldA string `json:"fieldA"`
		FieldB int    `json:"fieldB"`
	}
	reqtStc := &reqtStruct{
		FieldA: "bcd",
		FieldB: 2,
	}
	erro = redisCache.Store(ctxt, "test:test-struct", &reqtStc, 0)
	require.NoError(t, erro)
	respStc := &reqtStruct{}
	_, erro = redisCache.Retrieve(ctxt, "test:test-struct", &respStc)
	require.NoError(t, erro)
	require.Equal(t, reqtStc, respStc)
}
