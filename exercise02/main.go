package main 

import (
		"fmt"
		"sync"
)


func main(){
				fmt.Println("Exercise 2")
				// Объявляем массив со значениями {2, 4, 6, 8, 10}
				numbers := [5]int{2, 4, 6, 8, 10}
				//Использование и инициализация WaitGroup, механизма ожидания корутин
				wg := sync.WaitGroup{} 
				
				// Устанавливаем количество горутин, которые должны быть завершены
				wg.Add(len(numbers))
				for _, num := range numbers {
								// Запускаем горутину с текущим значением num
								go func(num int) {
												// Выводим квадрат числа
												fmt.Println(num * num)

												// Уменьшаем счетчик горутин в WaitGroup
												wg.Done()
								}(num) // Передаем num как аргумент, чтобы избежать захвата
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}	
						