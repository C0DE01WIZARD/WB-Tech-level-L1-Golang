// 25. Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"log"
	"time"
)

// Определяем функцию Sleep с определенной продолжительностью
func Sleep(valueDuration  time.Duration){
    // Создаём канал, который будет отправлен через определенное количество времени
    <-time.After(valueDuration)
}

func main(){
    fmt.Println("Exercise25")
    fmt.Printf("Установите время в секундах:") // Через консоль вводим время в секндах
    var valueDuration int // Определяем переменную valueDuration с типом: int для передачи значения в функцию
    fmt.Scan(&valueDuration) 

    sleepDuration := time.Duration(valueDuration) * time.Second // Преобразуем int в time.Duration
    Sleep(sleepDuration) // Передаем значение времени в секундах
    fmt.Printf("Программа продолжила работу через %d секунды после засыпания", valueDuration) // Получаем сообщение в консоль по работе функции Sleep
}