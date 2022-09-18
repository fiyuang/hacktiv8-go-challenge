package main

import (
	"fmt"
	"sync"
)

var (
	interface1 interface{}
	interface2 interface{}
	wg         sync.WaitGroup
)

func main() {
	interface1 = []string{"coba1", "coba2", "coba3"}
	interface2 = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go printProcessDisorder(i, &wg, interface1)
		go printProcessDisorder(i, &wg, interface2)
	}

	wg.Wait()
}

func printProcessDisorder(i int, wg *sync.WaitGroup, interfaceData interface{}) {
	fmt.Println(interfaceData, i)
	wg.Done()
}
