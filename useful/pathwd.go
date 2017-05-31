package useful

import (
	"fmt"
	"os"
	"runtime"
)

// GetPwd1 return .go file path, where func really is
//
// Useful only with .go file
// With binary file always return path of source .go file
func GetPwd1() string {
	_, pwd, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return pwd
}

// GetPwd2 always return folder path from where you run it
// For example if you run
// cd /folder
// then run in this path (we suppose that this file is exist)
// go run golang/src/project/main.go
// command return - /folder
func GetPwd2() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return pwd
}

// GetPwd3 return file path from where you run it (need binary file, not .go file)
// As of Go 1.8 (Released February 2017) the recommended way of doing PWD is with os.Executable
// But it's only useful when you run binary file (after go build)
// If you run it from .go file you will get something like
// /var/folders/n3/r0chsz09339gm2gbxdhjctgr0000gn/T/go-build824104956/command-line-arguments/_obj/exe/test
func GetPwd3() string {
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	return pwd
}
