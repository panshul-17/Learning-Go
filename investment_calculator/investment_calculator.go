package main

import (
	"fmt"
	"math"
)

// go is staticly typed language so types are important !
// go doesn't allow mixing types
// We have to do type conversions sometimes

func main() {

	const inflationRate = 2.5        // can't be reasigned
	var investmentAmount float64 = 0 // can be reassigned
	var years float64 = 0
	expectedReturnRate := 0.0
	fmt.Print("Enter the Amount")
	fmt.Scan(&investmentAmount) // expects a pointer to the variable
	fmt.Print("Enter the Amount and years")
	fmt.Scan(&years, &expectedReturnRate)
	value := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realValue := value / math.Pow(1+inflationRate/100, years)
	fmt.Println(value, realValue)

}
