package calc

import "github.com/myminicommission/api/graph/model"

// CommissionTotal takes in a commission and calculates the total
func CommissionTotal(commission *model.Commission) float64 {
	total := 0.00

	// add up the cost for all minis
	for _, mini := range commission.Minis {
		qty := float64(mini.Quantity)
		cost := mini.Price
		total = total + (qty * cost)

		// TODO: add up the options and multiply them for each mini
		// for _, option := range mini.Options {
		// 	total = total + (qty * option.Cost)
		// }
	}

	return total
}
