package main

import "fmt"

//import

// создаём структуру Human
type Human struct{
	 Name string
	 Age string
	 height int
}

// сначала объявляем дополнительный аргумент "(h Human)", далее описание описание функции
func (h Human) Speak(){
				fmt.Println("Hello, i am can speak!", h.Name)
}

// объявляем дополнительный метод Sleep
func (h Human) Sleep(){
				fmt.Println("Hello, i am can sleep!", h.Name)
}

type Action struct {}


