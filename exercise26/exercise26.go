/* 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например: 
abcd — true
abCdefAaf — false
aabcd — false */

package main

import (
    "fmt"
    "strings"
)

// Определяем функцию uniqueChars которая принимает str и возвращает false или true в зависимости от условий
func uniqueChars(str string) bool {
    // Создаём map для подсчета количества вхождений символов
    charCounter := make(map[rune]int) // Тип данных: rune. Используется для работы с символами для обработки каждого символа по отдельности
    for _, char := range str { // Перебираем все символы строки через цикл for
        charCounter[char]++ // Для каждого символа map счётчик увеличиваем на единицу
    }

    // Если количество вхождений каждого символа равно 1, то все символы уникальны
    for _, occurrenceСounter := range charCounter {
        if occurrenceСounter > 1 { // Проверяется условие, если хотя бы для одного символа количество вхождений больше 1, функция возвращает false
            return false // Если условие не сработало то возвращается false
        }
    }
    return true // Если условие сработало то возвращается true
}

func main() {
	fmt.Println("Exercise26")
    // Пример использования функции
    var inputChar string
    fmt.Print("Введите строку на проверку: ")
    fmt.Scan(&inputChar)
	// Условия при которых uniqueChars проверяет, содержит ли строка уникальные символы
    if uniqueChars(strings.ToLower(inputChar)) { //Функция strings.ToLower используется для приведения символа к нижнему регистру перед проверкой уникальности
        fmt.Println("Cимволы уникальны")
    } else {
        fmt.Println("Cимволы не уникальны")
    }
}