package main

import "fmt"

// OldLaptop - структура, представляющая старый ноутбук
type OldLaptop struct{}

// TurnOnPower - метод включения ноутбука в старом ноутбуке
func (OldLaptop) TurnOnPower() {
	fmt.Println("Включение питания ноутбука...")
}

// NewLaptop - структура, представляющая новый ноутбук
type NewLaptop struct{}

// UseBluetooth - метод, который использует Блютус в новом ноутбуке
func (NewLaptop) UseBluetooth() {
	fmt.Println("Поиск устройств Блютус...")
}

// UseWiFi - метод, который использует Wi-Fi в новом ноутбуке
func (NewLaptop) UseWiFi() {
	fmt.Println("Поиск Wi-Fi соединений...")
}

// LaptopAdapter - структура, которая служит адаптером между OldLaptop и NewLaptop
type LaptopAdapter struct {
	newLaptop NewLaptop
}

// TurnOnPower метод адаптера, который адаптирует вызов метода TurnOnPower старого ноутбука.
func (adapter *LaptopAdapter) TurnOnPower() {
	adapter.newLaptop.UseBluetooth() // Вызываем метод UseBluetooth нового ноутбука для добавления в старый ноутбук
}

func main() {
	fmt.Println("Exercise21")
	// Создаём новый ноутбук
	newLaptop := NewLaptop{}

	// Создаём адаптер для старого ноутбука
	oldLaptop := &LaptopAdapter{newLaptop: newLaptop}

	// Используем адаптер для старого ноутбука
	oldLaptop.TurnOnPower() // Включаем старый ноутбук
}
