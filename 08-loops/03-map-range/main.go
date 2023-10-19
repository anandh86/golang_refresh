package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(input map[int]string) []int  {

	inputKeys := []int{}

	for key := range input {
		inputKeys = append(inputKeys, key)
	}

	return inputKeys
}

func Values(input map[int]string) []string  {
	inputVals := []string{}

	for _,val := range input {
		inputVals = append(inputVals, val)
	}

	return inputVals

}
