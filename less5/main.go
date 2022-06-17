package main

import (
	"fmt"
	"log"
	"sync"
)

var (
	ctrlStack   []struct{}
	count, iter int
	mutex       sync.Mutex
	wg          = sync.WaitGroup{}
)

func main() {
	fmt.Print("количество потоков: ")
	fmt.Scanln(&count)

	wg.Add(count)
	for iter = 0; iter < count; iter++ {
		go func() {
			mutex.Lock()
			defer func() {
				mutex.Unlock()
				wg.Done()
			}()
			ctrlStack = append(ctrlStack, struct{}{})
		}()
	}
	wg.Wait()

	log.Printf("Количество отработавших потоков %d", len(ctrlStack))
}
