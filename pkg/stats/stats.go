package stats

<<<<<<< HEAD
import (
	"github.com/Shah-2021/bank/pkg/types"
)
=======
import "github.com/Shah-2021/bank/v2/pkg/bank/types"
>>>>>>> v2

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avgSum := types.Money(0)
	ind := types.Money(0)
<<<<<<< HEAD
	for _, v := range payments {
		avgSum += v.Amount
		ind ++
=======
	status := types.StatusFail
	for _, v := range payments {
		if v.Status != status  {
		avgSum += v.Amount
		ind ++
		}
>>>>>>> v2
	}
	return types.Money(avgSum/ind)
}


// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	totalSum := types.Money(0)
	for _, v := range payments {
<<<<<<< HEAD
		if v.Category == category {
			totalSum += v.Amount
		}
	}
=======
		if category == v.Category {
		if v.Status != types.StatusFail {
	    	totalSum += v.Amount
		}
	}
	}
>>>>>>> v2
	return totalSum
}