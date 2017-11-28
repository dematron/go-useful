package useful

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

//var prettyJSON bytes.Buffer
//err = json.Indent(&prettyJSON, []byte(data), "", "    ")
//if err != nil {
//	fmt.Println(err)
//}
//

type Iss struct {
	Securities struct {
		Columns []string        `json:"columns"`
		Data    [][]interface{} `json:"data"`
	} `json:"securities"`
}

var myTransport = &http.Transport{
	DisableCompression:    true,
	DisableKeepAlives:     true,
	ResponseHeaderTimeout: 3 * time.Second,
}
var myClient = &http.Client{Timeout: 3 * time.Second, Transport: myTransport}

const HOST = "https://moex.com"
const http_json = "?iss.json=compact&iss.meta=off"

func GetJsonAlt(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64;) Trident/7.0 (Touch; like Gecko) rv/11.0")
	r, err := myClient.Do(req)
	if err != nil {
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
			r, err := myClient.Do(req)
			if i == 3 {
				fmt.Println(err)
				con_test(url)
				return err
			} else if err != nil {
				time.Sleep(1000)
				continue
			} else {
				return json.NewDecoder(r.Body).Decode(target)
			}
		}
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println(err)
		con_test(url)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func JsonParsing() {
	iss := Iss{}
	iss_alt := Iss{}
	path := "/iss/engines/stock/markets/bonds/securities.json"
	URL := HOST + path + http_json
	GetJson(URL, &iss)
	fmt.Println("Run GetJson", iss)
	GetJsonAlt(URL, &iss_alt)
	fmt.Println("Run GetJsonAlt", iss_alt)
}

func JsonWithoutStruct() {
	value := gjson.Get(string(ReadUrl()), "full_name")
	fmt.Println(value.String())
}
