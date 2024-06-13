// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// Создаём структуру Struct_Counter с методами для инкрементирования и получения значения.
type Struct_Counter struct{
			value int // Объявляем в структуре Value с типом: int
			mutex sync.Mutex // Объявляем в структуре Mutex для синхронизации доступа к счётчику для избежания гонки данных.
}

// Функция для создания экземпляра структуры Struct_Counter, с начальным значение 0.
func CreateCounter() *Struct_Counter {
    return &Struct_Counter{value: 0} // Начальное значение Value:0
}


// Функция с методом Increment для инкрементирует значение счетчика.
func (counter *Struct_Counter) Increment() { // Метод Increment инкрементирует значение счетчика
    counter.mutex.Lock() // Блокируем мутекс для защиты от гонки данных (чтобы гарантировать, что только одна горутина может изменять значения в данный момент времени).
    defer counter.mutex.Unlock() // После изменения значения произойдёт разблокировка мутекса при помощи defer.
    counter.value++ // Увеличивает заначени Value на единицу
}

func (counter *Struct_Counter) Value() int { // Метод Value возвращает текущее значение счетчика
    counter.mutex.Lock() // Блокируем мутекс для защиты от гонки данных.
    defer counter.mutex.Unlock() // Разблокировка мутекса при ошибках вернуть значение счетчика.
    return counter.value
}

func main() {
		 fmt.Printf("Exercise 18\n")
		 counter := CreateCounter() // Создаём экземпляр структуры Struct_Counter

		 // Запускаем 50 горутин в цикле
		 for i := 0; i < 50; i ++ {
				go counter.Increment() // Запускаем 50 горутин, каждая из которых вызывает метод Increment
		 }
		 
		  // Канал для ожидания завершения всех горутин
    var done = make(chan struct{})

     // Закрываем канал после завершения всех горутин
    go func() {
        defer close(done) // Вызов defer и закрытие канала done
        for i := 0; i < 50; i++ { // Создаем горутину из 50 итераций
            counter.Increment()
        }
    }()

    // Ожидаем завершения всех горутин
    <-done

    // Итоговое значение счетчика.
    fmt.Println("Итоговое значение счетчика = ", counter.Value())
}
