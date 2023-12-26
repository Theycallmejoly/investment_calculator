package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var ExpectedReturnRate float64
	var years int

	fmt.Print("investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&ExpectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futurevalue := float64(investmentAmount) * math.Pow((1+ExpectedReturnRate/100), float64(years))
	futureRealValue := futurevalue / math.Pow(1+inflationRate/100, float64(years))

	formatedFV := fmt.Sprintf("the future value: %.2f \n", futurevalue)
	formatedRFV := fmt.Sprintf("the future value (adjusted for inflation): %.2f \n ", futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}
