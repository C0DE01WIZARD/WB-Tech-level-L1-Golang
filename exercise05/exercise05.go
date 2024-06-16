// 5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import ("fmt"
		"time")



func main(){
	fmt.Println("Exercise05")
	// Создание каналов
	channelData := make(chan int) // Создаем канал channelData для передачи данных 
	channelDone := make(chan bool)// Создаём канал channelDone, о сигнализировании завершения программы
	
	go sendingDataChannel(channelData) // Запускаем горутину sendingDataChannel для отправки данных в канал
    go readingDataChannel(channelData, channelDone)  // Запускаем горутину readingDataChannel для чтения данных из канала
    time.Sleep(5 * time.Second) // Ждем 5 секунд перед завершением программы
    channelDone <- true // Отправляем сигнал true о завершении в канал
}

// Инициализируем функцию sendingDataChannel, которая последовательно отправляет значения в канал
func sendingDataChannel(channelData chan int) { // В качестве аргументов передаём channelData с типом данных: int
    for ints := 0; ints < 5; ints++ { // Используем цикл for для последовательной отправки чисел от 0 до 4 в канал
        channelData <- ints // Отправляем значения ints в канал
        fmt.Printf("Отправили значение в канал: %d\n", ints)
        time.Sleep(1 * time.Second ) // Создаём паузу между отправками
    }
    close(channelData) // Закрываем канал, когда все значения отправлены
}

// Создаём функцию readingDataChannel, которая читает данные из канала и выводит их
func readingDataChannel(channelData chan int, channelDone chan bool) { // В качестве аргументов передаём channelData-для передачи данных в канал, channelDone-канал для сигнализирования о завершении работы функции
    for { // В бесконечном цикле for ожидаем данные на канале channelData или channelDone
        select { // Выбор условий
        case valueChannelData, condition := <-channelData:
            if !condition { // Проверка с помощью булевого значения, если condition=False, то канал закроется
                fmt.Println("Канал закрыт, завершаем чтение.") // И выводим в консоль сообщение о завершени чтения
                return // Завершаем чтение, если канал закрыт
            }

            fmt.Printf("Получили значение из канала: %d\n", valueChannelData) // Выводим значение полученное из канала
        case <-channelDone: // Ждем сигнала о завершении программы
            fmt.Println("Получили сигнал о завершении, завершаем чтение.") // Выводим сигнал о завершении
            return // Завершаем чтение при получении сигнала
        }
    }
}


