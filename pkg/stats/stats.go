package stats

import (
	"github.com/Shah-2021/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avgSum := types.Money(0)
	ind := types.Money(0)
	for _, v := range payments {
		avgSum += v.Amount
		ind ++
	}
	return types.Money(avgSum/ind)
}


// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	totalSum := types.Money(0)
	for _, v := range payments {
		if v.Category == category {
			totalSum += v.Amount
		}
	}
	return totalSum
}