package useful

import (
	"go-useful/common"

	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Read entire file at once, return []byte, but can string
func ReadEntireFile(path string) string {
	f, err := ioutil.ReadFile(path)
	common.CheckError(err)
	return string(f)
}

// Read file by lines
func ReadFileByLines(path string) {
	f, err := os.Open(path)
	common.CheckError(err)
	defer f.Close()

	n := 0
	scanner := bufio.NewScanner(f)
	//scanner.Split(bufio.ScanLines) used by default
	for scanner.Scan() {
		fmt.Println(n, scanner.Text())
		n++
	}
}
