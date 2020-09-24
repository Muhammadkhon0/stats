package stats

import (
	"github.com/Muhammadkhon0/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var mid types.Money
	for _, payment := range payments {
		if payment.Status == types.StatusOk{
			mid += payment.Amount
		}
	}
	length := len(payments)
	mid /= types.Money(length)
	return mid
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category == category && payment.Status == types.StatusOk{
			total += payment.Amount
		}
	}
	return total
}
