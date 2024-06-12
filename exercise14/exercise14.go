// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import ("fmt"
				"reflect" //  импорт пакета для работы с типами и значениями во время выполнения
)


func main(){
		fmt.Println("Exercise14")
		// Объявляем слайс, который содержит: число, строку, булево значение и канал.
		type_variable :=[]interface{}{123, "String", true, make(chan int)}


		// Пробегаемся по слайсу type_variable
		for _, type_variables := range type_variable{
				// Используем `reflect.TypeOf` для получения типа переменной type_variables для каждого типа данных
				fmt.Println("Тип данных в слайсе определен как >", reflect.TypeOf(type_variables))
		}
}
