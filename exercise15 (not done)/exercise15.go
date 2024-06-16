/* 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

// Ответ: Данный код приводит к проблемам с производительностью.

package main

import "fmt"

func main() {
  fmt.Println("Exercise15")

}