// 7. Реализовать конкурентную запись данных в map.


package main

import ("fmt"
		"sync" // Импорт пакета для синхронизации горутин
)


var mutex sync.Mutex // Объявляем переменную mutex типа sync.Mutex для синхронизации блокировки доступа к ресурса. Предотврашение гонки данных
var dataMap = make((map[int]string)) // Объявляем переменную dataMap тип: map[int]string для конкурентной записи данных.

// Создаём функцию writerKey с аргументом keyMap с типом int
func writerKey(keyMap int){
	mutex.Lock() // Блокируем доступ к dataMap.!!! Только одна горутина может изменять map !!!
	defer mutex.Unlock() // Производим разблокировку после выполнения функции
	dataMap[keyMap] = "Запись в dataMap" // Запись ключа и значения в dataMap
}


func main() {
	fmt.Println("Exercise07")
	var waitgroup sync.WaitGroup // Создаём переменную waitgroup с типом данных: sync.WaitGroup - для управления ожиданиями горутинами
	waitgroup.Add(10) // Добавляем 10 горутин в ожидание на выполнение с помощью Add 

	for argumentKey :=0; argumentKey < 10; argumentKey++ { // Запускаем счетчик на for 10 горутин каждая из которых вызывает функцию writerKey с аргументом argumentKey
		go func (argumentKey int){ // Создаем анонимную функцию с аргументом argumentKey с типом: int
			defer waitgroup.Done() // После выполнения уменьшаем счетчик горутин
			writerKey(argumentKey) // Выполняем запись в dataMap

		}(argumentKey) // Аргумент, который передаётся в функцию writerKey
	}


waitgroup.Wait() // Ожидание завершения всех горутин

  for keyMap, valueMap := range dataMap { // Перебираем данные в dataMap.
        fmt.Printf("Ключ: %d, Значение: %s\n", keyMap, valueMap) // Выводим в консоль ключи и значения.
    }
}