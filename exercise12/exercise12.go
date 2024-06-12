// 12.Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

func main(){
		fmt.Println("Exercise 12")
		// Создаем слайс stringsList с значениями из задания
		stringsList := []string {"cat", "cat", "dog", "cat", "tree"}

		// Создаем слайс для хранения уникальных значений
    uniqueSlice := make([]string, 0)

		// Пробегаемся по слайсу с заданными значениями
    
		for _, valueSlice := range stringsList {
					 // Используем флаг для отслеживания, было ли значение уже добавлено в слайс
        flag := true
        
				// Пробегаемся по слайсу с уникальными значениями
        for _, val := range uniqueSlice {
            
					// Проверяем, есть ли значение в слайсе
            if valueSlice == val {
                flag = false
            }
        }
 // Если значение не найдено в слайсе, добавляем его
        if flag {
            uniqueSlice = append(uniqueSlice, valueSlice)
        }
    }

    // Выводим слайс с уникальными значениями
    fmt.Printf("Уникальный набор элементов множества: %v\n", uniqueSlice)

}