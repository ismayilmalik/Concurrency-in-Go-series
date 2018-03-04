package code3

import (
	"sync"
	"fmt"
)

func main() {
	wg:= sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("second goroutine")
	}()
	
	wg.Wait()
	fmt.Println("main goroutine")
}

