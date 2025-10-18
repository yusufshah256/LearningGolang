// Common fmt verbs:
// %v - default format
// %T - type of the value
// %d - decimal integer
// %f - floating-point number
// %s - string
// %t - boolean
// %q - quoted string
// %x/%X - hex encoding (lower/upper case)
// %b - binary
// %c - character
// %o/%O - octal
// %e/%E - scientific notation (lower/upper case)
// %U - Unicode format
// %% - literal percent sign
// \n - new line
// \t - tab
// \r - carriage return
// \b - backspace
// \f - form feed
// \v - vertical tab
// \\ - backslash
// ' - single quote
// \" - double quote
// %d - decimal integer or bases 10

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	var name string
	var number int
	var age int
	var floatNum float64
	var isActive bool
	var isDone bool

	fmt.Printf("Enter your Name: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter you Age: ")
	fmt.Scanln(&age)

	fmt.Printf("Enter your integer Number: ")
	fmt.Scanln(&number)

	fmt.Printf("Enter your Float Number: ")
	fmt.Scanln(&floatNum)

	fmt.Printf("Are you Active (true/false): ")
	fmt.Scanln(&isActive)

	fmt.Printf("Are you Done (true/false): ")
	fmt.Scanln(&isDone)

	result := fmt.Sprintf("Hello, %v and you are %d years old and your request is being processed please wait...", name, age)
	io.WriteString(os.Stdout, result)

	// Pause for 10 seconds
	time.Sleep(10 * time.Second)

	// Integer related verbs: %v, %d, %b, %c, %o, %O, %x, %X

	fmt.Printf(
		"\nThe integer number is %v\n"+
			"and it's corresponding value in the base 16 with upper case is %X\n"+
			"and it's corresponding value in the base 16 with lower case is %x\n"+
			"and it's base 2 is %b\n"+
			"and it's corresponding unicode char is %c (unicode point U+%[5]X)\n"+
			"and it' base 10 is %d\n"+
			"and it's base 8 is %o\n"+
			"and it's unicode format is %U\n"+
			"and it's pointer is %p\n"+
			"and it's quoted value is %q\n"+
			"and it's base 8 with 0o prefix is %O\n\n",
		number, number, number, number, number, number, number, number, &number, fmt.Sprintf("%d", number), number,
	)

	// String related verbs: %s, %q, %x, %X
	fmt.Printf(
		"The string is %s\n"+
			"and it's value is %v\n"+
			"and its quoted string is %q\n"+
			"and it's base 16 lower case is %x\n"+
			"and it's base 16 upper case is %X\n\n",
		name, name, name, name, name,
	)

	// Float related verbs: %v, %f, %e, %E
	fmt.Printf(
		"The float number is %v\n"+
			"and it's value with 2 decimal points is %.2f\n"+
			"and it's value with 4 decimal points is %.4f\n"+
			"and it's scientific notation with lower case e is %e\n"+
			"and it's scientific notation with upper case E is %E\n\n",
		floatNum, floatNum, floatNum, floatNum, floatNum,
	)

	// Boolean related verbs: %v, %t
	fmt.Printf("The boolean is %t and it's values is %v the type is %T\n", isActive, isActive, isActive)
	fmt.Printf("The boolean is %t and it's value is %v the type is %T\n", isDone, isDone, isDone)

}
