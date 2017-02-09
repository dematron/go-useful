package useful

import (
	"go-useful/common"

	"os"
	"path/filepath"
	"strings"
)

// Strings
// Join
func PathJoin() string {
	pwd, err := os.Getwd()
	common.CheckError(err)
	result := filepath.Join(pwd, "//some//other/path without/_errors")
	return result
}

// Join
func StringsJoinOne() string {
	pwd, err := os.Getwd()
	common.CheckError(err)
	result := (pwd + "//some//other/path with/_errors")
	return result
}

// Join
func StringsJoinTwo() string {
	pwd, err := os.Getwd()
	common.CheckError(err)
	s := []string{"\n/foo", "\n/bar", "\n/baz"}
	result := strings.Join(s, pwd)
	return result
}

// Split
func StringsSplit() []string {
	pwd, err := os.Getwd()
	common.CheckError(err)
	result := strings.Split(pwd, "/")
	return result
}
