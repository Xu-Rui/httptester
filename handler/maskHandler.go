package handler

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomMask(kvmap map[string]string) {
	for k, v := range kvmap {
		if strings.Contains(k, "time") {
			break
		}
		value, err := strconv.Atoi(v)
		if err == nil {
			kvmap[k] = strconv.Itoa(value + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10))
		} else {
			kvmap[k] = createRandomStr()
		}
	}
	if _, ok := kvmap["mask"]; ok {
		kvmap["mask"] = "4"
	}
	if _, ok := kvmap["uid"]; ok {
		kvmap["uid"] = "256"
	}
}

func createRandomStr() string {
	nums := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < 6 {
		num := byte(r.Intn(60) + 60)
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return string(nums)
}
