package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- fmt.Sprintf("Воркер %d обработал задачу %d", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
