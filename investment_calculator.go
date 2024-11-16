package main

import "fmt"

func main(){
	var investmentAmount = 1000;
	var expectedReturnRate = 5.5;
	var years = 10;

	var futureValue = (((investmentAmount/100)*int(expectedReturnRate))*years)+investmentAmount
	fmt.Print(futureValue)
}