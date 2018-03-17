package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func TestGenerateNewSessionID(t *testing.T) {
	session, err := GenerateNewSessionID();
	assert.Empty(t, err)
	fmt.Printf("session id is : %s", session)
	assert.NotEmpty(t, session,)
}

func TestRedisGetNil(t *testing.T) {
	conn := GetRedisClient().Get()
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("GET", "test"))
	assert.Equal(t, redis.ErrNil, err)
	assert.Empty(t, b)
}

func TestLoadYamlConfigFile(t *testing.T) {
	LoadYamflConfigFile("../config.yml")
}