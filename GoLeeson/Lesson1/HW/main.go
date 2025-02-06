package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		parts := strings.Split(job, " ")

		if len(parts) != 3 {
			results <- "Неверный формат задачи"
			continue
		}

		oper := parts[0]
		num1, _ := strconv.Atoi(parts[1])
		num2, _ := strconv.Atoi(parts[2])
		var result int

		if oper == "+" {
			result = num1 + num2
		} else if oper == "-" {
			result = num1 - num2
		} else {
			results <- "Неверный оператор"
			continue
		}

		results <- fmt.Sprintf("Результат %s: %d", job, result)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 5)
	results := make(chan string, 5)

	// Добавляем задачи в канал jobs
	jobs <- "+ 5 3"
	jobs <- "- 10 4"
	close(jobs)

	// Запускаем воркеров
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Ожидаем завершения всех воркеров
	wg.Wait()
	close(results) // Закрываем канал results после завершения всех воркеров

	// Выводим результаты из results
	for result := range results {
		fmt.Println(result)
	}
}
