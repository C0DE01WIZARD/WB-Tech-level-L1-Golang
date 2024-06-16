// 17. Реализовать бинарный поиск встроенными методами языка.

package main 

import "fmt" 

// Создаём функцию binarySearch выполняет бинарный поиск элемента element в отсортированном срезе values
func binarySearch(values []int, element int) int {
	lostartingIndex := 0 // Определяем начальный индекс среза
	endingIndex := len(values) - 1 // Определяем конечный индекс среза

	// Цикл продолжается, пока не будет проверен весь диапазон
	for lostartingIndex <= endingIndex {
		midRange := (lostartingIndex + endingIndex) / 2 // Находим середину диапазона
		midValue := values[midRange] // Значение в середине диапазона

		// Сравниваем midValue с искомым элементом e
		if midValue < element {
			lostartingIndex = midRange + 1 // Сдвигаем начальный индекс за середину
		} else if midValue > element {
			endingIndex = midRange - 1 // Сдвигаем конечный индекс перед середину
		} else {
			return midRange // Нашли элемент, возвращаем его индекс
		}
	}

	return -(lostartingIndex + 1) // Элемент не найден, возвращаем отрицательное значение
}

func main() {
	fmt.Println("Exercise17")
	nums := []int{2, 4, 6, 8, 10} // Отсортированный срез чисел для поиска
	fmt.Println("Число '6' индекс элемента",binarySearch(nums, 6), "в слайсе") // Выводим  в консоль результат поиска числа 6
	fmt.Println("Число '10' индекс элемента",binarySearch(nums, 10), "в слайсе") // Выводим в консоль результат поиска числа 10
	fmt.Println("Число '2' индекс элемента",binarySearch(nums, 2), "в слайсе") // Выводим в консоль результат поиска числа 2
	fmt.Println("Число '8' индекс элемента",binarySearch(nums, 8), "в слайсе") // Выводим в консоль результат поиска числа 8
	fmt.Println("Число '4' индекс элемента",binarySearch(nums, 4), "в слайсе") // Выводим в консоль результат поиска числа 4
}
