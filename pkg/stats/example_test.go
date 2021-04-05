package stats

import (
	"fmt"

<<<<<<< HEAD
	"github.com/Shah-2021/bank/pkg/types")


func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 100,
		},
		{
			Amount: 80,
		},
		{
			Amount: 99,
		},
	}

    fmt.Println(Avg(payments))

	//Output: 93
=======
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

>>>>>>> v2
}
