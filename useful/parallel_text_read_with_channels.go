package useful

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
)

func matchTxt(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	n := 0

	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for j := range jobs {
		fmt.Println("retry", n)
		fmt.Println(j)
		n++
		if strings.Contains(j, "7") {
			results <- j
		}
	}
}

// goOverFile - go over a file line by line and queue up a ton of work
func goOverFile(f io.Reader, jobs chan<- string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Later I want to create a buffer of lines, not just line-by-line here ...
		jobs <- scanner.Text()
	}
	close(jobs)
}

// TxtFile - Main function
func TxtFile(txt string) string {
	f := strings.NewReader(txt)

	// do I need buffered channels here?
	jobs := make(chan string)
	results := make(chan string)

	// I think we need a wait group, not sure.
	wg := new(sync.WaitGroup)

	// start up some workers that will block and wait?
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go matchTxt(jobs, results, wg)
	}

	go goOverFile(f, jobs)

	// Now collect all the results...
	// But first, make sure we close the result channel when everything was processed
	go func() {
		wg.Wait()
		close(results)
	}()

	result := ""
	for i := range results {
		fmt.Println("Catch:", i)
	}

	return result
}

//func main() {
//	// An artificial input source.  Normally this is a file passed on the command line.
//	//const input = "Foo\n(555) 123-3456\nBar\nBaz\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456"
//
//	path := "/etc/os-release"
//	if _, err := os.Stat(path); !os.IsNotExist(err) {
//
//		f, err := ioutil.ReadFile(path)
//		common.CheckError(err)
//
//		n := TxtFile(string(f))
//		fmt.Println(n)
//	}
//}
