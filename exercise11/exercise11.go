// 11.Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func main(){
		fmt.Println("Exercise 11")
		// Создаём два слайса для представления множеств
		set1 :=make([]int, 0) // 1 множество
		set2 :=make([]int, 0) // 2 множество

		// добавляем значения в множества. 
		//Примечание: добавляем одинаковые значения в двух множествах чтоб реализовать перечение, согласно заданию.

		//Добавляем значения с помощью функции append в первое множество set1
		set1 = append(set1, 1)
		set1 = append(set1, 2)
		set1 = append(set1, 3)

		//Добавляем значения с помощью функции append во второе множество set2
		set2 = append(set2, 4)
		set2 = append(set2, 2)
		set2 = append(set2, 3)

		// Создаем map для хранения. Ключи: int, значения: bool
    sliceMap := make(map[int]bool)

		 // Пробегаемся по set1 и устанавливаем флаги для каждого значения
    for _, valueSlice1 := range set1 {
        sliceMap[valueSlice1] = true
    }

		 // Пробегаемся по set2 и устанавливаем флаги для каждого значения
    for _, valueSlice2 := range set2 {
        if sliceMap[valueSlice2] {
							fmt.Printf("Число %d - общее вхождение в обоих множествах\n", valueSlice2)
				}
    }


	}

