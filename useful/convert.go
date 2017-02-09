package useful

import (
	"go-useful/common"

	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Convert slice ([]int) to string
func ConvertSliceToString(request interface{}) (result string) {
	//fmt.Println(reflect.TypeOf(request).String())
	if reflect.TypeOf(request).String() == "[]int" {
		// x.(T)
		// asserts that x is not nil and that the value stored in x is of type T.
		// The notation x.(T) is called a type assertion.
		r := request.([]int)
		//r, ok := request.([]int) // Alt. non panicking version
		result = convertIntSliceToString(r)
	} else if reflect.TypeOf(request).String() == "[]string" {
		r := request.([]string)
		// Join our string slice.
		result = strings.Join(r, " ")
	} else {
		result = strings.Trim(fmt.Sprint(request), "[]")
	}
	return result
}

func convertIntSliceToString(request []int) (result string) {
	valuesText := []string{}
	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range request {
		number := request[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	// Join our string slice.
	result = strings.Join(valuesText, " ")
	return result
}

// Atoi (string to int) and Itoa (int to string)
func ConvertAtoiItoa() (string, int) {
	integer := int(32)
	result1 := strconv.Itoa(integer)
	result2, err := strconv.Atoi(result1)
	common.CheckError(err)

	return result1, result2
}

// Convert String to []byte
func StringToByteSlice(content string) []byte {
	return []byte(content)
}
