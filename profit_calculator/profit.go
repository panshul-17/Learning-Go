package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("enter revenue, expenses and taxRate")
	fmt.Scan(&revenue, &expenses, &taxrate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax / 100 * (100 - taxrate)
	ratio := earningBeforeTax / earningAfterTax

	fmt.Print("EBT ", earningBeforeTax, " Proft ", earningAfterTax, " Ratio", ratio)

}
