package stats

import (
	"fmt"

	"github.com/Shah-2021/bank/v2/pkg/bank/types"
)


func ExampleAvg() {
	payments := []types.Payment {
		{
			Amount: 100,
			Status: types.StatusOk,
			Category: "аптеки",
		},
		{
			Amount: 80,
			Status: types.StatusInProgress,
			Category: "аптеки",
		},
		{
			Amount: 99,
			Status: types.StatusOk,
			Category: "рестораны",
		},
		{
			Amount: 99,
			Status: types.StatusFail,
			Category: "аптеки",
		},
	}


fmt.Println(Avg(payments))
fmt.Println(TotalInCategory(payments, "аптеки" ))
	//Output: 
	// 93
	// 180

}
