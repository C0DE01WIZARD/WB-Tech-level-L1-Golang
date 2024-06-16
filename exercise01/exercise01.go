package main
import "fmt"

// Cоздаём родительскую структуру Human
type Human struct{
	 Name string
	 Age int
	 height int
}

// Cначала объявляем дополнительный аргумент "(human Human)", далее описание описание функции Speak
func (human *Human) Speak(){
				fmt.Println("Hello, i am can speak!", human.Name)
}

// Объявляем метод Sleep
func (human *Human) Sleep(){
				fmt.Println("Hello, i am can sleep!", human.Name)
}

// Объявляем метод Height
func (human *Human) Height(){
				fmt.Println("Hello, my height", human.height)
}

// Создаём Action - дочерную структуру
type Action struct {
				Human // Реализация встраивания Human в стркутуру Action
}

func Exercise_1(){
			fmt.Printf("Exercise 1\n")
			fmt.Printf("Human структура - родительская:\n")
		
			// Создание объекта Human
			person := Human{"Dmitry", 22, 185}
			
			// Вызов методов
			person.Sleep()
			person.Speak()
			person.Height()

			fmt.Printf("\nAction структура - дочерняя :\n")
			// Создание объекта Action и встраивание Human в него
			action := Action{person}
			// Вызов методов
			action.Sleep()
			action.Speak()
			action.Height()
}
func main() {

		Exercise_1()
}