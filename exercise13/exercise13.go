// 13. Поменять местами два числа без создания временной переменной

package main

import "fmt"

// Объявляем функцию swap_variable для реализации обмена переменных a,b
func swap_variable(a int,b int) (int, int){
			a,b =b,a
			return a,b
}


func main(){
		fmt.Println("Exercise 12")
		a := 1
		b := 2

		// Меняем местами переменные a,b
		a,b = swap_variable(a,b)
		// Выводим значений после замены
		fmt.Println("Значения перемнных после замены", "a =",a,"b =",b)

}