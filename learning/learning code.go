package main

import "fmt"

func main() {
    fmt.Printf("\n========== Exercise 10: ==========\n\n")

    // Создаем карту для вывода и слайс с заданными значениями
    temps := make(map[int]([]float32))
    given_vars := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

    // Перебираем слайс и добавляем значения в карту
    for _, value := range given_vars {
        // int(value) / 10 * 10 - округление до ближайшего десятка
        temps[int(value)/10*10] = append(temps[int(value)/10*10], value)
    }

    fmt.Printf("Temperatures: %v\n", temps)

    fmt.Printf("\n====================================\n\n")
}