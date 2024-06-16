package main 

import (
		"fmt"
		"sync"
)


func main(){
			fmt.Println("Exercise 02")
			// Объявляем массив со значениями {2, 4, 6, 8, 10}
			numbers := []int{2, 4, 6, 8, 10}
			//Использование и инициализация WaitGroup, механизма ожидания корутин
			wg := sync.WaitGroup{} 
			
			// Устанавливаем количество горутин, которые должны быть завершены и передаём длину слайса
			//Мы запускаем столько горутин, сколько элементов, содержится в слайсе numbers
			wg.Add(len(numbers))
			for _, number := range numbers {
							
				// Запускаем горутину с текущим значением number
			go func(number int) {
							// Выводим квадрат числа
							square := number * number
							fmt.Printf("Квадрат числа %d равен %d\n", number, square)
							
							// Уменьшаем счетчик горутин в WaitGroup по завершению
							wg.Done()
			}(number) // Передаем number как аргумент, чтобы избежать захвата
}

	// Ожидаем завершения всех горутин
	wg.Wait()
}	
						