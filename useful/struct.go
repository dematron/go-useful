package useful

import (
	"go-useful/common"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Struct declaration
type TestStruct struct {
	One   int
	Two   string
	Three []string
}

// Using structs
func UsingStruct() (result *TestStruct) {
	slice := []string{"0", "*", "*", "*", "*", "root", "sh", "/usr/local/bin/toggle_swap_.sh"}
	fmt.Println(slice)
	one, err := strconv.Atoi(slice[0])
	common.CheckError(err)
	result = &TestStruct{
		One:   one,
		Two:   slice[1],
		Three: slice[2:],
	}

	return result
}

// if struct in bufio reading file (.NewScanner and .Scan) you should use schema like
// .Scan -> .Text -> populate struct like popStruct := &MyStruct{ -> at the end - var = append(var, popStruct)
// and return - var with type *MyStruct or []*MyStruct
// example in cron.go
//

// TODO: https://play.golang.org/p/Pw9f20zwja

// Configuration
type Configuration struct {

}

func getConfig(path string, c interface{}) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
	}
}

func GetMainConfig(path string) Configuration {
	c := Configuration{}
	getConfig(path, &c)
	return c
}
