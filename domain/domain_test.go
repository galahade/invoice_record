package domain

import (
	"testing"
	"fmt"
	"time"
	_"github.com/json-iterator/go"
)

type testJsonTime struct {
	message string
	date JsonTime
}

func TestJsonTimeUnmarshal(t *testing.T) {
	testJson := "\"2018-03-12\""
	fmt.Printf("test json is : %s\n", testJson)
	jsonTime := new(JsonTime)
	jsonTime.UnmarshalJSON([]byte(testJson))
	fmt.Printf("UnmarshalJSON's JsonTime is %s\n", time.Time(*jsonTime))
}

func TestJsonTimeMarshall(t *testing.T) {
	testValue := JsonTime(time.Now())
	fmt.Printf("testValue is %#v\n", testValue)
	b, err := testValue.MarshalJSON()
	fmt.Printf("error is : %s\n", err)
	fmt.Printf("marshl result is %s\n", string(b))
}
/*
func TestJsonTimeUnmarshalJSON(t *testing.T) {
	testJson := "{\"message\":\"test json time unmarshal\",\"date\":\"2018-03-12\"}"
	fmt.Printf("test json is : %s\n", testJson)
	testValue := new(testJsonTime)
	jsoniter.Unmarshal([]byte(testJson), testValue)
	fmt.Printf("UnmarshalJSON's message is %s\n", testValue.message)
	fmt.Printf("UnmarshalJSON's JsonTime is %s\n", time.Time(testValue.date))
}

func TestJsonTimeMarshallJSON(t *testing.T) {
	testValue := new(testJsonTime)
	testValue.date = JsonTime(time.Now())
	testValue.message = "test json time marshal"
	fmt.Printf("testValue is %#v\n", testValue)
	b, err := jsoniter.Marshal(testValue)
	fmt.Printf("error is : %s\n", err)
	fmt.Printf("marshl result is %s\n", string(b))
}
*/