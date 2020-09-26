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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	moneys := map[types.Category]types.Money{}
	quantity := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			moneys[payment.Category] += payment.Amount
			quantity[payment.Category]++
		}
	}

	for key := range moneys {
		moneys[key] /= quantity[key]
	}
	return moneys
}
