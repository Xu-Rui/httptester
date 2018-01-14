package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	R_Finish  = "RealtimeFinishEvent"
	R_Already = "RealtimeAlreadyGrab"
)

var ParamsMap map[string]string

func init() {

	ParamsMap := make(map[string]string)

	b, err := ioutil.ReadFile("testcase")
	if err != nil {
		fmt.Print(err)
	}
	str := strings.Split(strings.Trim(string(b), " "), "#####")

	for i := 0; i < len(str); {
		if i+2 < len(str) {
			ParamsMap[str[i+1]] = strings.Trim(str[i+2], " ")
		}
		i += 3
	}

	fmt.Println(ParamsMap[R_Finish])
}

func testPost(Params string) (string, error) {
	var request *gorequest.SuperAgent
	request = gorequest.New().Post("http://localhost:8803/gulfstream/nereus/collectEvent")
	request.Set("didi-header-rid", "123")
	request.Set("didi-header-spanid", "123")
	request.Set("didi-header-hint-code", "")
	request.Set("didi-header-hint-content", "")
	request.Set("Content-Type", "application/json")
	data := make(map[string]string)
	data["Params"] = ParamsMap[Params]
	data["caller"] = "kafka-consumer"
	_, body, err := request.Type("json").Send(data).EndBytes()
	if len(err) != 0 {
		return "", err[0]
	}

	result := strings.Replace(strings.Trim(strings.Trim(string(body), "{"), "}"), ",", "\n", -1)
	result = "\n---" + Params + "---\n" + result + "\n-------------------------\n"
	return result, nil
}

func main() {
	fmt.Println(testPost(R_Finish))
}
