package code1

import (
	"fmt"
)

func main() {
	go printHelloWorld()
}

func printHelloWorld() {
	fmt.Println("Hello World")
}
