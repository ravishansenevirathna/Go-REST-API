package examples

import (
	"fmt"
	"math"
)

func Invest() {

	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = (((investmentAmount / 100) * int(expectedReturnRate)) * years) + investmentAmount
	fmt.Println(futureValue)
	//power of
	fmt.Println(math.Pow(2, 3))

	var x int = 5
	var y float64 = 99
	fmt.Println("x is :", x)
	fmt.Println("y is :", y)

	//go referred variable declaration
	// this is the go recommended shorter way to assign a variable
	testMarks := 85
	fmt.Println("your mark is :", testMarks)

	const yearTaxRate = 5.2
	fmt.Println("constance is defined here", yearTaxRate)

	//input and output
	//In Go, you must pass a pointer (using the & operator) to the fmt.Scan()
	//function because it needs to modify the value of the variable where the input is stored.

	var username string
	print("enter your name : ")
	fmt.Scan(&username)
	fmt.Println("your name is : ", username)

	// get the two user input as number and print the maximum number.

	var num1 int
	var num2 int
	fmt.Print("Enter your First number : ")
	fmt.Scan(&num1)
	fmt.Print("Enter your Second number : ")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println("Maximum Number is :", num1)
	} else {
		fmt.Println("Maximum Number is :", num2)
	}

	//function example
	myFunction(25)

}

func myFunction(num int) {
	fmt.Println("Hello from myFunction", num)
}
