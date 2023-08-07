package main

import (
	"fmt"

	"sync"
	"time"

	pkg "m16/pkg"
)

func main() {

	var mu sync.Mutex
	value := 3

	fmt.Println("we start")

	go pkg.Accumulate(value)
	mu.Lock()
	pkg.Action()
	mu.Unlock()

	time.Sleep(time.Second)

}
