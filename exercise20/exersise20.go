// 20. Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
// !!! Задание похоже с 19-ым.!!! 

package main

import (
    "fmt"
    "strings" // импорт модуля strings для работы со строками
)

// Функция для переворачивания строки
func reverseString(s string) string {
    runes := []rune(s) 	// Преобразуем строку в срез Unicode-символов (runes) 
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Переворачиваем срез по аналогии в задании 19
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes) // Возвращаем строку, полученную из среза runes
}

func main() {
    fmt.Println("Exercise20")
    string_input := "snow dog sun" // Определяем переменную string_input со значение "snow dog sun", согласно задания
    reversed := reverseWords(string_input) // Переворачиваем слова в строке
    fmt.Println(reversed)  // Выводим результат
}

// Функция для переворачивания слов в строке
func reverseWords(s string) string {
    words := strings.Split(s, " ") // Разделяем строку на слова
    reversedWords := make([]string, len(words)) // Переворачиваем каждое слово
    for i, word := range words {
        reversedWords[i] = reverseString(word)
    }
    return strings.Join(reversedWords, " ") // Соединяем слова обратно в строку
}