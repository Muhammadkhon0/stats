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
	money := map[types.Category]types.Money{}
	quantity := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			money[payment.Category] += payment.Amount
			quantity[payment.Category]++
		}
	}

	for key := range money {
		money[key] /= quantity[key]
	}
	return money
}

func PeriodsDynamic(first, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for key, amount := range second {
		result[key] += amount
	}
	for key, amount := range first {
		result[key] -= amount
	}
	return result
}

