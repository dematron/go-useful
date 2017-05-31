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

	// Different 'pwd' variants
	fmt.Println("Path of the the file where func GetPwd1() declared:", useful.GetPwd1())
	fmt.Println("Running golang file folder path:", useful.GetPwd2())
	fmt.Println("Running golang binary file path:", useful.GetPwd3())
	pwd := useful.GetPwd2()
	file := "LICENSE"
	pth := pwd + "/" + file
	fmt.Println("Path to needed file LICENSE:", pth)

	// parallel text read with channels
	fmt.Println()
	fmt.Println("Parallel read from LICENSE with channels")
	if _, err := os.Stat(pth); !os.IsNotExist(err) {

		f, err := ioutil.ReadFile(pth)
		common.CheckError(err)

		n := useful.TxtFile(string(f))
		fmt.Println(n)
	}

	// Url read
	fmt.Println()
	fmt.Println("Simple URL content read example")
	fmt.Println(string(useful.ReadUrl()))

	// Using json parsing from http page
	fmt.Println()
	fmt.Println("Get json content from http GET")
	useful.JsonParsing()

	// Json parsing without struct
	fmt.Println()
	fmt.Println("Read JSON without struct")
	useful.JsonWithoutStruct()
}
