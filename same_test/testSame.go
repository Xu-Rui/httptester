package same_test

import (
	"encoding/json"
	"fmt"
	"github.com/Xu-Rui/httptester/dao"
	"github.com/parnurzeal/gorequest"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	oldURL = "http://0.0.0.0/test01"
	newURL = "http://0.0.0.0/test02"
)

func TestSame(basecase ICase) {

	//获得测试数据集合
	res := basecase.GetTestCases()

	for i, v := range res {

		//获取一条测试数据
		datamap := basecase.Mockdata(v)

		//40次/1s
		time.Sleep(time.Duration(50) * time.Millisecond)
		oldResult := basecase.Test(datamap, oldURL)
		newResult := basecase.Test(datamap, newURL)
		result := reflect.DeepEqual(oldResult, newResult)

		if !result {
			fmt.Println()
			fmt.Println("第", i, "组")
			fmt.Println("old: ", oldResult)
			fmt.Println("new: ", newResult)
			fmt.Println("params: ", datamap)
		}

		if (i+1)%1000 == 0 {
			fmt.Println(" ", i+1)
			fmt.Print("=")
		} else if (i+1)%200 == 0 {
			fmt.Print("=", " ", i+1, " ")
		} else if (i+1)%10 == 0 {
			fmt.Print("=")
		}

	}

}
