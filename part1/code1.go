package part1

import (
	"fmt"
)

func main() {
	go printHelloWorld()
}

func printHelloWorld() {
	fmt.Println("Hello World")
}
