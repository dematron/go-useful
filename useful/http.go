package useful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

var myTransport = &http.Transport{
	DisableCompression:    true,
	DisableKeepAlives:     true,
	ResponseHeaderTimeout: 3 * time.Second,
}
var myClient = &http.Client{Timeout: 3 * time.Second, Transport: myTransport}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

func con_test(url string) {
	var body []byte
	var response *http.Response
	var request *http.Request

	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		debug(httputil.DumpRequestOut(request, true))
		response, err = (&http.Client{}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		debug(httputil.DumpResponse(response, true))
		body, err = ioutil.ReadAll(response.Body)
	}

	if err == nil {
		fmt.Printf("%s", body)
	} else {
		log.Fatalf("ERROR: %s", err)
	}
}

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
