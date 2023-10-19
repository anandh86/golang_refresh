package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(args ...int) int  {

	sum := 0

	for _, val := range args {
		sum += val
	}

	return sum

}
