package main

import (
	"fmt"
	"time"
)

func main() {
	Mylang := "Mylang"
	fmt.Println("Hello, World!", Mylang)
	time.Sleep(time.Second * 9)
	Mylang2 := Mylang + "2"
	fmt.Println("Hello, World!", Mylang2)

}
