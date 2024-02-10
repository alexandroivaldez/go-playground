package main

import (
	"fmt"
	"math"
)

func main() {
	// Go Variables
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	// Typical casting like in Java and other strong type langauges
	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println(futureValue)
}
