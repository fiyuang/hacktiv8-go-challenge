package main

import (
	"fmt"
	"sync"
)

var (
	interface1 interface{}
	interface2 interface{}
	wg         sync.WaitGroup
	m          sync.Mutex
)

func main() {
	interface1 = []string{"coba1", "coba2", "coba3"}
	interface2 = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		m.Lock()
		go printProcess(i, &wg, &m, interface1)
		m.Lock()
		go printProcess(i, &wg, &m, interface2)
	}

	wg.Wait()
}

func printProcess(i int, wg *sync.WaitGroup, m *sync.Mutex, interfaceData interface{}) {
	fmt.Println(interfaceData, i)
	m.Unlock()
	wg.Done()
}
