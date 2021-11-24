package main

import "fmt"

func p(nums map[string]int) {
	fmt.Println("nums = {")
	for key, value := range nums {
		fmt.Printf("  %s -> %d\n", key, value)
	}
	fmt.Println("}")
}

func main() {
	var nums map[string]int = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	p(nums)

	fmt.Println("nums.len =", len(nums))
	fmt.Println("nums =", nums)

	nums["four"] = 4
	fmt.Println("nums.len =", len(nums))
	fmt.Println("nums =", nums)

	fmt.Println("nums[\"zero\"] =", nums["zero"])
	fmt.Println("nums[\"two\"] =", nums["two"])
	fmt.Println("nums[\"argh\"] =", nums["argh"])

	delete(nums, "zero")
	delete(nums, "three")
	fmt.Println("nums.len =", len(nums))
	fmt.Println("nums =", nums)
	fmt.Println("nums[\"zero\"] =", nums["zero"])
	fmt.Println("nums[\"one\"] =", nums["one"])
	fmt.Println("nums[\"two\"] =", nums["two"])
	fmt.Println("nums[\"three\"] =", nums["three"])
	fmt.Println("nums[\"four\"] =", nums["four"])

	p(nums)
}
