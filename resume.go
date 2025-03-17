package main

import "fmt"
import "sync"

func main() {
    current_lang := "Perl"
    var profit float64

    if employerOffersAwesomeTasks(current_lang) {
        profit = workHardWith(current_lang);
    } else {
        profit = take_profit_from_cpp_and_golang();
    }

    fmt.Printf("Your profit will be: %.2f ₽\n", profit)
}

func take_profit_from_cpp_and_golang() (profit float64) {
    langs := []string{ "C++", "Golang" }
    var wg sync.WaitGroup
    var mu sync.Mutex

    for priority, lang := range(langs) {
        wg.Add(1);
        go func(lang string, priority int) {
            defer wg.Done()
            setGoroutinePriorityTo(priority)
            partial_profit := workHardWith(lang)
            mu.Lock()
            defer mu.Unlock()
            profit += partial_profit
        }(lang, priority);
    }

    wg.Wait();

    return
}

func employerOffersAwesomeTasks(lang string) bool {
    // Your implementation here
}

func workHardWith(lang string) (profit float64) {
    // TODO
}

func setGoroutinePriorityTo(value int) {
    // Incredible implementation here
}
