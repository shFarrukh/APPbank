package stats

import (
	"fmt"

	"github.com/shFarrukh/bank/pkg/bank/types")


func ExampleAvg() {
	peyments := []types.Payment {
		{ 
			Amount : 10_000,
		},
		{
			Amount:  20_000,
		},
	}
	result := Avg(peyments)
	fmt.Println(result)
	// Output
	// 15000
}