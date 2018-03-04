package code3

import (
	"fmt"
)

func main() {
	go printTime()
	fmt.Println("main goroutine")
}

func printTime() {
	fmt.Println("second goroutine")
}