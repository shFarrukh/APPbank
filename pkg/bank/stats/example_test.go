package stats

import (
	"fmt"

	"github.com/shFarrukh/bank/pkg/bank/types")


func ExampleAvg() {
	peyments := []types.Payment {
		{ 
			Amount : 10_000,
			Category: "car",
		},
		{
			Amount:  20_000,
			Category: "car",
		},
	}
	result := Avg(peyments)
	fmt.Println(result)
	// Output
	// 15000
}