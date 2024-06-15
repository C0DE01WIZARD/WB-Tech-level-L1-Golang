// 23.Удалить i-ый элемент из слайса.

package main

import ("fmt")

func main(){
	fmt.Println("Exercise23")
	// Создаем слайс со строковым типом из 4 элементов 
	// !!! индексы от 0 до 3, [0]="Golang", [1]="super", [2]="Powerful", [3]="Language"
	sliceString := []string {"Golang", "super", "powerful", "language"}
	fmt.Printf("Введите номер индекса для удаления: ") // Вводим в консоль номер индекса
    var indexSelection int // Объявляем переменную indexSelectionс со строковым типом string
    fmt.Scan(&indexSelection) // Считываем индекс.

	// Проверка переменной indexSelection на диапазон нахождения введенного индекса от 0 до 3. Проверка необходима для избежания ошибок во время выполнения программы.
	if indexSelection < 0 || indexSelection >= len(sliceString) { // !!! ||- логический оператор OR("ИЛИ") для выполнения двух булевых выражений 
        fmt.Println("Индекса не существует, ничего не удалено")
    }  else {
        // Используем метод append для создания нового слайса, который состоит из всех элементов исходного слайса до индекса включительно
        sliceString = append(sliceString[:indexSelection], sliceString[indexSelection+1:]...)
        fmt.Printf("Новый слайс после удаления i-го элемента: %v\n", sliceString) // Выводим в консоль новый слайс после удаления i-го элемента
    }
}