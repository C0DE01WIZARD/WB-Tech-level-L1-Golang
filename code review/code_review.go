package main

import (
    "fmt"
    "strings"
)

var justString string

func createHugeString(size int) string {
    // Предположим, что createHugeString создает строку заданного размера.
    return strings.Repeat("a", size)
}

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100] // Используем только часть строки
    v = "" // Освобождаем память, освобождая v
}

func main() {
    someFunc()
    fmt.Println("justString:", justString)
}