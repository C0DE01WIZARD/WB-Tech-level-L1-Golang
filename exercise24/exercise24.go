// 24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import ( "fmt"
		"math"
)
// Создаём структуру с двумя полями x и y-координаты точки с типом float64
type Point struct{
	x,y float64
}

// Создаём конструктор 'createPoint', который возвращает новую точку с координатами.
func createPoint(x, y float64) *Point {
	return &Point{x, y}
}


func square(num float64) float64{
	return num*num
}

// Создаём функцию Distance для возвращения значения расстояния между двумя точками. 
// !!! Для результата вычислений будем использовать тип: float64 для предоставления наилучшей точности и производительности !!!
func Distance(point1, point2 *Point) float64 {
	// Используем теорему Пифагора для вычисления расстояния
	// !!! Для удобства разделили формулу D = √(x₂ - x₁) ² + (y₂ - y₁)²  на несколько частей
    dx := point1.x - point2.x // Вычисляем расстояние, вычитая point1.x(x1) - point2.x(x2)
    dy := point1.y - point2.y // После этого мы должны вычесть point1.y(y1) - point2.y(y2)
	resulstSquareX := square(dx)
	resulstSquareY := square(dy)
	allresult := resulstSquareX+resulstSquareY
    return math.Sqrt(allresult) // Далее возводим dx и dy в квадрат и складываем

	/* Вариант №2 как проше посчитать по теореме Пифагора без создания дополнительной функции square 
    return math.Sqrt(dx*dx + dy*dy) // Далее возводим dx и dy в квадрат и складываем */  
}

func main(){
	fmt.Println("Exercise24")
	// Создаем две точки point1 point2 для вычисления
    point1 := createPoint(6, 3)
    point2 := createPoint(3, 6)
   
	// Вычисляем и выводим расстояние между точками
    distance := Distance(point1, point2)
    fmt.Printf("Расстояние между точками: %.2f\n", distance) // Выводим в консоль формате %.2f указывает на то, что число должно быть выведено с двумя знаками после запятой
}