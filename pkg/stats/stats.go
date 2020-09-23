package stats

import (
	"github.com/Muhammadkhon0/bank/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	var mid types.Money
	for _, payment := range payments {
		mid += payment.Amount
	}
	length := len(payments)
	mid /= types.Money(length)
	return mid
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}
	return total
}
