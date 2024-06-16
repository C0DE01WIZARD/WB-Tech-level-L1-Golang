/* 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/* Ответы: 
1)Данный код приводит к проблемам с производительностью памяти такой как "утечка". Выражение 1 << 10 в переменной v c методом createHugeString - означает сдвиг влево числа  1 на 10 бит. В десятичной системе будет число больше 1024 байта (2*10=1024), но извлекается подстрока длиной всего 100 символов v[:100].

2) Переменная v будет содержать ссылку на строки, которая очень большая и только часть будет использоваться в justString.

3) В переменной v память не будет освобождена сборщиком мусора, что может привести к утечке памяти.

Решения:
1) Использование константы const 1024 вместо 1 << 10 - для явного указания значения без битовых операций
2) Не имеет смысл хранить огромную строку без явного использования или необходимости, приводящая в использованию больших ресурсов памяти. 
Необходимо использовать оптимизированные структуры данных или методы, которые не требуют хранения всей строки в памяти сразу. 
3) Переменную v присваивоим пустой строке, чтобы освободить память, занимаемую исходной строкой.

Реализация:
*/

package main

import ("fmt"
         "strings" // импорт модуля для работы со строками
        )

var justString string

// Создаём функцию createHugeString для создания строки заданного размера
func createHugeString(size int) string{ // Добавляем аргумент size с типом данных: int и возвращаем string
    return strings.Repeat("a", size) // Возвращает заданное количество раз повторение строки через strings.Repeat
}

// Создаем функцию someFunc для вызова createHugeString для создания строки размером 1024 символа
func someFunc(){
    v := createHugeString(1 << 10) // Создание строки размером 1024 символа
    justString = v[:100] // Сохраняем в переменную justString первые 100 символов
    v = " " // Освобождаем память, через сборщика мусора
}


func main() {
   fmt.Println("Exercise15")
   someFunc() // Вызываем функцию someFunc
   fmt.Println("JustString:", justString) // Выводим в консоль значения переменной justString
}
