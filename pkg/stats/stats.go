package stats

import (
	"github.com/Muhammadkhon0/bank/v2/pkg/types"
)


func Avg(payments []types.Payment) types.Money {
	count := len(payments)
	sum := 0

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += int(payment.Amount)
	}
	avg := sum / (count - 1)

	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}