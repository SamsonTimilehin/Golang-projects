
// Simple Calculator with Golang

package main
import(
	"fmt"
)

var symbol string
	var num1 int
	var num2 int
	var option string
	var totalResult int
	   
	
	


func main(){

fmt.Println("Hey Welcome")

    
	
	collectUserInput()
   calculate(num1,num2,symbol)
   continutionFunction()

}

func calculate(number1 int, number2 int, operator string){
	
		
	var result int	
	
	if operator == "+" {
		result = number1 + number2
		totalResult += result
	
	}else if operator == "-" {
		result = number1 - number2
		totalResult -= result
	
	}else if operator == "*" {
		result = number1 * number2
		totalResult *= result
	
	}else if operator == "/" {
	
		if number2 == 0 {
			fmt.Print(" second value cannot be zero", number2)
		}else {
	
		result = number1 / number2
		totalResult /= result
	
		}
	}else {
		fmt.Print("Invalid operator\n")
	}
	
	fmt.Printf("The result is %v\n",totalResult)
	
}

func collectUserInput(){

	fmt.Println("Enter first number")
	fmt.Scan(&num1)
	
	fmt.Println("Enter operator")
	fmt.Scan(&symbol)
	
	fmt.Println("Enter second number")
	fmt.Scan(&num2)

	fmt.Println("Press y to continue or n to stop")
	fmt.Scan(&option)	


}

func continutionFunction(){
	for{

	if option == "y"{
	  collectUserInput2()
	  calculate(num1,num2,symbol)
	}else if option == "n"{
		break
	}
}
}

func collectUserInput2(){

	fmt.Println("Enter operator")
	fmt.Scan(&symbol)

	fmt.Println("Enter first number")
	fmt.Scan(&num1)

	fmt.Println("Enter operator")
	fmt.Scan(&symbol)
	
	fmt.Println("Enter second number")
	fmt.Scan(&num2)

	fmt.Println("Will you like to continue the calculation")
	fmt.Scan(&option)
	
}