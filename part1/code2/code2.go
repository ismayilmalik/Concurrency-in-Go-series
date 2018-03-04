package code2

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Hello World")
	}()
}
