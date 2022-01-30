
// Simple Calculator with Golang

package main
import(
	"fmt"
)

var symbol string
	var num1 int
	var num2 int


func main(){

fmt.Println("Hey Welcome")

    
	
	collectUserInput()
   calculate(num1,num2,symbol)

}

func calculate(number1 int, number2 int, operator string){
	
	
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

func collectUserInput(){

	fmt.Println("Enter first number")
	fmt.Scan(&num1)
	
	fmt.Println("Enter operator")
	fmt.Scan(&symbol)
	
	fmt.Println("Enter second number")
	fmt.Scan(&num2)
}