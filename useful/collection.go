package useful

import (
	"strings"
	"fmt"
)

func Rename(strs []string, f func(string) string) []string {
	var mstrs []string
	for _, v := range strs {
		v = f(v)
		mstrs = append(mstrs, v)
	}
	return mstrs
}

func CollectionExample() {
	strs := []string{"one", "two", "five"}
	fmt.Println(Rename(strs, strings.Title))
}
