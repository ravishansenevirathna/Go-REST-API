package main

import "fmt"
import "math"

func main(){
	var investmentAmount = 1000;
	var expectedReturnRate = 5.5;
	var years = 10;

	var futureValue = (((investmentAmount/100)*int(expectedReturnRate))*years)+investmentAmount
	fmt.Println(futureValue);
	//power of 
	fmt.Println(math.Pow(2,3));

	var x int =5;
	var y float64 = 99;
	println("x is :",x);
	println("y is :",y);


	//go refered variable declaration
	// this is the go recommended shorter way to assign a variable
	testMarks := 85
	println("your mark is :",testMarks);

}