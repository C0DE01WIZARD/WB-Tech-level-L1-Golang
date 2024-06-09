package main

import (
	"fmt"
	"time"
)

func main(){
		go fmt.Println("Hello goroutine!!!")
		go fmt.Println("Hello from main()")
		time.Sleep(time.Microsecond)

}