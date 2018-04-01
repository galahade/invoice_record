package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func TestGenerateNewSessionID(t *testing.T) {
	session:= GenerateNewSessionID();
	fmt.Printf("session id is : %s", session)
	assert.NotEmpty(t, session,)
}

func TestRedisGetNil(t *testing.T) {
	pool := GetRedisPool(LoadYamlConfigFile(GetRootPath()+"/config.yml"))
	defer pool.Close()
	conn := pool.Get()
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("GET", "test"))
	assert.Equal(t, redis.ErrNil, err)
	assert.Empty(t, b)
}


func TestGetRootPath(t *testing.T) {
	path := GetRootPath()
	log.Printf("Root Path is : %s", path)
}



func TestTestEnv(t *testing.T) {
	testEnv()
}