package main

import (
	"fmt"
)

func main() {
	const myConst string = "This is the Constant String var"
	A := "This a Constant String Var"
	fmt.Println(A)
	A = "This is a Changed String Var"
	fmt.Println(A)
	fmt.Println(myConst)
}
