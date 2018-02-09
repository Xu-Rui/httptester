package dao

import (
	"fmt"
	redisClient "github.com/Xu-Rui/httptester/dao/redis"
	"testing"
)

func init() {
	err := redisClient.InitCodis("0.0.0.0:6379")
	if err != nil {
		fmt.Println(err)
	}
}

func Test_InitCodis(t *testing.T) {
	err := redisClient.InitCodis("0.0.0.0:6379")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func Test_Get(t *testing.T) {
	result, err := redisClient.RedisClient.Get("COT_25")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(result)
}

func Test_Set(t *testing.T) {
	err := redisClient.RedisClient.Set("name", "xurui")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	result, err := redisClient.RedisClient.Get("name")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if result != "xurui" {
		t.Errorf("result not correct||get=%v||should be=%v", result, "xurui")
		t.Fail()
	}
}
