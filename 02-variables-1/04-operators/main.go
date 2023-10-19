package main

import "fmt"

func main() {
	numberOfProducts := 3
	productCost := 100
	shippingCost := 10

	totalCost := productCost * numberOfProducts
	totalCost += shippingCost

	fmt.Println("Total order cost:", totalCost)
}
