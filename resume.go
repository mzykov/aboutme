package main

import (
	"fmt"
	"sync"
)

func main() {
	langs := []string{"C++", "Golang"}

	var profit float64
	var wg sync.WaitGroup
	var mu sync.Mutex

	for priority, lang := range langs {
		wg.Add(1)
		go func(lang string, priority int) {
			defer wg.Done()
			setGoroutinePriorityTo(priority)
			partial_profit := workHardWith(lang)
			mu.Lock()
			defer mu.Unlock()
			profit += partial_profit
		}(lang, priority)
	}

	wg.Wait()

	fmt.Printf("Your profit will be: %.2f ₽\n", profit)
}

func workHardWith(lang string) (profit float64) {
	// Implementation here
	return 100000000
}

func setGoroutinePriorityTo(value int) {
	// Incredible implementation here
}
