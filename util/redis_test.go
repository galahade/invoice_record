package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetRedisClient(t *testing.T) {
	cfg := LoadYamflConfigFile("../config.yml")
	pool := GetRedisClient(cfg)
	conn := pool.Get()
	result, err := conn.Do("SET", "key", "value", "100")
	result, err = conn.Do("EXPIRE", "key", 100)
	if err != nil {
		fmt.Printf("set key error is %s", err)
	}
	assert.Empty(t, err)
	fmt.Printf("set key - value to redis result is %s", result)
}

func TestRedisKeys(t *testing.T) {
	cfg := LoadYamflConfigFile("../config.yml")
	conn := GetRedisClient(cfg).Get()
	var err error
	if b, err1 := conn.Do("KEYS", "invoice::o5gGe4khB5GaEXO-Dn2waDD13zSs::*"); err1 == nil {
		assert.NotEmpty(t, b)
		fmt.Printf("KEYS command results are : %s", b)
	} else {
		err = err1
		fmt.Printf("KEYS error is : %s", err)
	}
	assert.Empty(t, err)
}

func TestRedisKeysNoResult(t *testing.T) {
	cfg := LoadYamflConfigFile("../config.yml")
	conn := GetRedisClient(cfg).Get()
	var err error
	if b, err1 := conn.Do("KEYS", "invoice::o5gGe4khB5GaEXO::*"); err1 == nil {
		assert.Empty(t, b)
		fmt.Printf("KEYS command results are : %s", b)
	} else {
		err = err1
		fmt.Printf("KEYS error is : %s", err)
	}

	assert.Empty(t, err)
}