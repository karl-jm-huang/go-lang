package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type item struct {
	Location   map[string]string `json:"location"`
	Now        map[string]string `json:"now"`
	LastUpdate string            `json:"last_update"`
}

type resultJSON struct {
	Results []item `json:"results"`
}

// HTTPGet .
func HTTPGet(q string) string {
	resp, err := http.Get(generateURL(q))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var resultjson resultJSON
	if err = json.Unmarshal(body, &resultjson); err == nil {
		return fmt.Sprint(
			"\n城市: "+resultjson.Results[0].Location["name"],
			"\n国家: "+resultjson.Results[0].Location["country"],
			"\n天气现象: "+resultjson.Results[0].Now["text"],
			"\n露点温度: "+resultjson.Results[0].Now["temperature"],
			"\n数据更新时间: "+resultjson.Results[0].LastUpdate,
			"\n")
	}

	return q + ": Ocurr error when parsing json response"
}

// HTTPGetAsync .
func HTTPGetAsync(q string, s chan string) {
	resp, err := http.Get(generateURL(q))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var resultjson resultJSON
	if err = json.Unmarshal(body, &resultjson); err == nil {
		s <- fmt.Sprintln(
			"\n城市: "+resultjson.Results[0].Location["name"],
			"\n国家: "+resultjson.Results[0].Location["country"],
			"\n天气现象: "+resultjson.Results[0].Now["text"],
			"\n露点温度: "+resultjson.Results[0].Now["temperature"],
			"\n数据更新时间: "+resultjson.Results[0].LastUpdate)
	} else {
		fmt.Println(q + ": 解析json出错")
		s <- ""
	}
}

func generateURL(q string) string {
	return "https://api.seniverse.com/v3/weather/now.json?key=zcsw0aalioc5lyf8&location=" + q + "&language=zh-Hans&unit=c"
}
