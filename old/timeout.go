package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"

	}()

	// fmt.Println("current c1:", c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case test := <-time.After(1 * time.Second):
		fmt.Println("timeout 1", test)
	}

}
