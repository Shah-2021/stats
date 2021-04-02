package stats

import (
	"fmt"

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
}
