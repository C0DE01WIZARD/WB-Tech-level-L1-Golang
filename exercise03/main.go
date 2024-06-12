package main

import (
	"fmt"
	"sync"
)

// Функция square вычисляет квадрат числа и отправляет данные вычислений в канал squares.
func square(wg *sync.WaitGroup, numbers <-chan int, squaeres chan<- int ){
		for num := range numbers { // Пробегаемся по каждому числу из канала numbers
				squaeres <- num *num // После вычисления квадрата числа данные отправляем в канал squares
		}
		wg.Done() // Завершаем горутину
}


func main(){
			numbers := [] int {2,4,6,8,10} // Инициализируем массив чисел {2,4,6,8,10}
			numberChan := make(chan int, len(numbers)) // Канал для чисел, равным размеру массива
			squaresChan := make(chan int, len(numbers)) // Канал для квадратов чисел
			var wg sync.WaitGroup // Ожидание синхронизации горутин

			// Вычисление квадратов
			for i :=0; i<len(numbers); i++{
				wg.Add(1) // Увеличиваем счётчик на единицу
				go square(&wg, numberChan, squaresChan) // Запуск горутины
			}
			
			//Отправка чисел в канал
			for _, number:= range numbers{
				numberChan <-number
			}
			close(numberChan)
			// Ожидание завершения всех горутин
    wg.Wait()
    close(squaresChan)
 // Суммирование квадратов
    sum := 0
    for square := range squaresChan {
        sum += square
    }

    fmt.Printf("Сумма квадратов чисел: %d\n", sum)
}



