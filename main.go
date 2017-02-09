package main

import (
	"go-useful/common"
	"go-useful/useful"

	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

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
	pwd, err := os.Getwd()
	common.CheckError(err)
	file := "LICENSE"
	path := pwd + "/" + file
	if _, err := os.Stat(path); !os.IsNotExist(err) {

		f, err := ioutil.ReadFile(path)
		common.CheckError(err)

		n := useful.TxtFile(string(f))
		fmt.Println(n)
	}
}
