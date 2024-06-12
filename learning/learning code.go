package main

import (
	"fmt"
	_ "fmt"
)

func main() {
  data := make(chan int)
  exit := make(chan int)
    go func()  {
      for i:=0; i<10; i++{
        fmt.Println(<-data)
      }
      exit <- 0
    }()
    selectOne(data,exit)
}

func selectOne(data, exit chan int){

}
