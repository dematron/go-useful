package main

import (
	"go-useful/common"
	"go-useful/useful"

	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Iss struct {
	Securities struct {
		Columns []string        `json:"columns"`
		Data    [][]interface{} `json:"data"`
	} `json:"securities"`
}

const HOST = "https://moex.com"
const http_json = "?iss.json=compact&iss.meta=off"

// Results
func main() {
	fmt.Println(useful.ReadEntireFile("/etc/hosts"))
	useful.ReadFileByLines("/etc/hosts")
	// Return variable type
	variable := float64(155.28)
	GetVarType := func() {
		fmt.Println("The type of the variable is:", reflect.TypeOf(variable))
	}
	GetVarType()
	//
	fmt.Println("Return slice:", useful.CreateSlice())
	//
	fmt.Println("Convert slice to string:", useful.ConvertSliceToString(useful.CreateSlice()))
	//
	fmt.Println("Return map:", useful.CreateMap())
	fmt.Println("Return map value by key \"5\" :", useful.CreateMap()[5])
	// Strings
	// Join
	fmt.Println(useful.PathJoin())
	fmt.Println(useful.StringsJoinOne())
	fmt.Println(useful.StringsJoinTwo())
	// Split
	fmt.Println(useful.StringsSplit())
	// Atoi and Iota convert
	a, b := useful.ConvertAtoiItoa()
	fmt.Println("First should be string:", a, reflect.TypeOf(a))
	fmt.Println("Second should be int:", b, reflect.TypeOf(b))
	// Struct
	str := useful.UsingStruct()
	fmt.Println("Struct, Three:", str.Three)
	// with pointers
	//c := useful.Dump()
	//for _, r := range c {
	//	//as func
	//	fmt.Println(useful.DumpSwapf(r))
	//	//as method
	//	fmt.Println(r.DumpSwaps())
	//}
	// String to []byte
	str1 := "Hello, world!"
	fmt.Println("String 'Hello, world!' ToByteSlice", useful.StringToByteSlice(str1))

	// parallel text read with channels
	fmt.Println()
	fmt.Println("Parallel read from LICENSE with channels")

	//EDIT: As of Go 1.8 (Released February 2017) the recommended way of doing PWD is with os.Executable
	//pwd, err := os.Executable()
	_, pwd, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	pwd = strings.Trim(pwd, "main.go")
	file := "LICENSE"
	pth := pwd + file
	fmt.Println(pth)
	if _, err := os.Stat(pth); !os.IsNotExist(err) {

		f, err := ioutil.ReadFile(pth)
		common.CheckError(err)

		n := useful.TxtFile(string(f))
		fmt.Println(n)
	}

	// Using json parsing from http page
	fmt.Println()
	fmt.Println("Get json content from http GET")

	iss := Iss{}
	iss_alt := Iss{}
	path := "/iss/engines/stock/markets/bonds/securities.json"
	URL := HOST + path + http_json
	useful.GetJson(URL, &iss)
	fmt.Println("Run GetJson", iss)
	useful.GetJsonAlt(URL, &iss_alt)
	fmt.Println("Run GetJsonAlt", iss_alt)
}
