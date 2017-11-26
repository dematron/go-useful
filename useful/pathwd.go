package useful

import (
	"go-useful/common"

	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetPwd1 return .go file path, where func really is
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
	common.CheckError(err)
	return pwd
}

// GetPwd3 return file path from where you run it (need binary file, not .go file)
// As of Go 1.8 (Released February 2017) the recommended way of doing PWD is with os.Executable
// But it's only useful when you run binary file (after go build)
// If you run it from .go file you will get something like
// /var/folders/n3/r0chsz09339gm2gbxdhjctgr0000gn/T/go-build824104956/command-line-arguments/_obj/exe/test
func GetPwd3() string {
	pwd, err := os.Executable()
	common.CheckError(err)
	return pwd
}

// GetPwdOther print 3 more variants of getting different paths
func GetPwdOther() {
	fmt.Printf("Binary file name with filepath.Base(os.Args[0]): %q.\n", filepath.Base(os.Args[0]))
	fmt.Printf("Full path to binary file including its name with filepath.Clean(os.Args[0]): %q.\n", filepath.Clean(os.Args[0]))
	fmt.Printf("Full path to binary file excluding its name with filepath.Dir(os.Args[0]): %q.\n", filepath.Dir(os.Args[0]))
}
