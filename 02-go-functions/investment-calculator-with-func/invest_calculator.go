package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	//fmt.Print("Invesment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	/*Version 1:*/

	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)

	//Example Output:
	//Future Value: 1628.766474448336
	//Future Value (adjusted for Inflation): 1272.766474448336

	/*Version 2:*/

	//fmt.Printf("Future Value: %v\nFuture Value (adjusted for Inflation): %v", futureValue, futureRealValue)

	//Example Output:
	//Future Value: 1628.766474448336
	//Future Value (adjusted for Inflation): 1272.766474448336

	/*Version 3:*/

	//fmt.Printf("Future Value: %.0f\nFuture Value (adjusted for Inflation): %.0f", futureValue, futureRealValue)

	//Example Output:
	//Future Value: 1628
	//Future Value (adjusted for Inflation): 1272

	/*Version 4:*/

	/*fmt.Printf(`Future Value: %.0f
	Future Value (adjusted for Inflation): %.0f`, futureValue, futureRealValue)*/

	//Example Output:
	//Future Value: 1629
	//	Future Value (adjusted for Inflation): 1272

	/*Version 5:*/

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

	//Example Output:
	//Future Value: 1628.9
	//Future Value (adjusted for Inflation): 1272.5
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
	//return
}
