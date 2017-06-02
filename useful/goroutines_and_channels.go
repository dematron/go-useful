package useful

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Channel1 Simple unbuf channel
func Channel1() {
	msg := make(chan int)

	go func() { msg <- 0110 }()

	m := <-msg
	fmt.Println(m)
}

// Channel2 Buf channel
func Channel2() {
	msg := make(chan string, 100)

	for i := 0; i < 10; i++ {
		msg <- "Message " + strconv.Itoa(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-msg)
	}
}

// Channel3 Goroutine control channel
func Channel3(done chan bool) {
	fmt.Print("Job started.")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("Job finished!")

	done <- true
}

// channel4_1 support only channel "c" for sending
func channel4_1(c chan<- string, msg string) {
	c <- msg
}

// channel4_2 support channel "c" for getting and "cc" for sending
func channel4_2(c <-chan string, cc chan<- string) {
	msg := <-c
	cc <- msg
}

// Channel4 Using of in-out channels
func Channel4() {
	p1 := make(chan string, 1)
	p2 := make(chan string, 1)
	channel4_1(p1, "Here can be your message.")
	channel4_2(p1, p2)

	fmt.Println(<-p2)
}

// Channel5 Example of two parallel channels working
func Channel5() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// channel6_1 is a worker for Channel6 which also add different sleep
func channel6_1(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	item := <-jobs
	if item == "a" {
		time.Sleep(5 * time.Second)
	} else if item == "b" {
		time.Sleep(3 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
	}
	results <- "result of " + item
}

// Channel6 show how to use channels in loop for similar work with different execution time
func Channel6() {
	start := time.Now()

	jobs := make(chan string)
	results := make(chan string)
	wg := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go channel6_1(jobs, results, wg)
	}

	for _, i := range []string{"a", "b", "c"} {
		if i == "b" {
			fmt.Println(i)
			jobs <- i
		}
		if i == "c" {
			fmt.Println(i)
			jobs <- i
		}
		if i == "a" {
			fmt.Println(i)
			jobs <- i
		}
	}

	for i := 0; i <= 2; i++ {
		fmt.Println(<-results)
	}

	// Close
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println(time.Since(start))
}

// goroutine1_1 is a worker for Goroutine1, similar to channel6_1
func goroutine1_1(item string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(item)
	if item == "a" || item == "d" {
		time.Sleep(5 * time.Second)
	} else if item == "b" || item == "e" {
		time.Sleep(3 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("result of " + item)
}

// Goroutine1 is similar to Channel6 but with using of goroutines and waitgroups
func Goroutine1() {
	start := time.Now()

	wg := new(sync.WaitGroup)

	for _, i := range []string{"a", "b", "c"} {
		if i == "b" {
			wg.Add(1)
			go goroutine1_1(i, wg)
		}
		if i == "c" {
			wg.Add(1)
			go goroutine1_1(i, wg)
		}
		if i == "a" {
			wg.Add(1)
			go goroutine1_1(i, wg)
		}
	}

	// Close
	wg.Wait()
	fmt.Println(time.Since(start))

	wg = new(sync.WaitGroup)

	for _, i := range []string{"d", "e", "f"} {
		if i != "" {
			wg.Add(1)
			go goroutine1_1(i, wg)
		}
	}

	// Close
	wg.Wait()
	fmt.Println(time.Since(start))
}

// goroutineAndChannels1_1 worker for GoroutineAndChannels1
func goroutineAndChannels1_1(item int, result chan<- []int) {
	if item == 1 {
		time.Sleep(5 * time.Second)
	} else {
		time.Sleep(2 * time.Second)
	}
	list := []int{1 + item, 2 + item, 3 + item}

	fmt.Println(list)

	result <- list
}

// GoroutineAndChannels1 is a combination of using goroutines and channels
func GoroutineAndChannels1() {
	start := time.Now()

	total1 := 0
	total1list := make(chan []int)
	total2 := 0
	total2list := make(chan []int)

	for _, i := range []string{"g", "h", "i"} {
		if i == "h" {
			go goroutineAndChannels1_1(1, total1list)
		}

		if i == "g" {
			go goroutineAndChannels1_1(2, total2list)
		}
	}

	//
	total1total1 := 0
	list := <-total1list
	for _, x := range list {
		total1total1 += x
	}
	total1 += total1total1 / len(list)

	//
	total2total2 := 0
	list = <-total2list
	for _, x := range list {
		total2total2 += x
	}
	total2 += total2total2 / len(list)

	fmt.Println(total2)
	fmt.Println(total1)

	fmt.Println(time.Since(start))
}

func GoroutineAndChannelsExamples() {
	//
	fmt.Println("Example of Channel1")
	Channel1()
	//
	fmt.Println("Example of Channel2")
	Channel2()
	//
	fmt.Println("Example of Channel3")
	// Declare chan
	done := make(chan bool, 1)
	//start
	go Channel3(done)
	//stop
	<-done
	//
	fmt.Println("Example of Channel4")
	Channel4()
	//
	fmt.Println("Example of Channel5")
	Channel5()
	//
	fmt.Println("Example of Channel6")
	Channel6()
	//
	fmt.Println("Example of Goroutine1")
	Goroutine1()

	//
	fmt.Println("Example of GoroutineAndChannels1")
	GoroutineAndChannels1()
}
