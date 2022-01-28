package main
import(
	"fmt"
)

// Simple Calculator with Golang

func main(){

fmt.Println("Hey Welcome")

var operator string
var number1 int
var number2 int

fmt.Println("Enter first number")
fmt.Scan(&number1)

fmt.Println("Enter operator")
fmt.Scan(&operator)

fmt.Println("Enter second number")
fmt.Scan(&number2)

var result int

if operator == "+" {
    result = number1 + number2


}else if operator == "-" {
	result = number1 - number2
	

}else if operator == "*" {
	result = number1 * number2
	

}else if operator == "/" {

	if number2 == 0 {
		fmt.Print(" second value cannot be zero", number2)
	}else {

	result = number1 / number2
	

	}
}else {
	fmt.Print("Invalid operator")
}

fmt.Printf("The result is %v",result)

}