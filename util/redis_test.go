package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	pool := GetRedisPool(LoadYamlConfigFile(GetRootPath() + "/config.yml"))
	defer pool.Close()
	conn := pool.Get()
	defer conn.Close()

	t.Run("Set expire key", func(t *testing.T) {
		_, err := conn.Do("SET", "key", "value", "100")
		result, err := conn.Do("EXPIRE", "key", 100)
		if err != nil {
			fmt.Printf("set key error is %s", err)
		}
		assert.Empty(t, err)
		fmt.Printf("set key - value to redis result is %s", result)
	})


}