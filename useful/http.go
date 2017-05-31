package useful

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

var url = "https://api.github.com/repos/dematron/go-useful"

func ReadUrl() []byte {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		// err
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//if resp.StatusCode == 200 { // OK
	//	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	//	bodyString := string(bodyBytes)
	//}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// err
		fmt.Println(err)
	}
	return bodyBytes
}

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
