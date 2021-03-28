package stats

import (
	"github.com/shFarrukh/bank/pkg/bank/types"
)


func Avg(payments []types.Payment) types.Money{
	s := 0
	couter := 0
	for _, i := range payments {

		s+= int(i.Amount)
		couter ++
	}
	return types.Money(s/couter)

}