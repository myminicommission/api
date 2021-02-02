package calc

import "github.com/myminicommission/api/graph/model"

// EstimateTotal takes in an estimate and calculates the total
func EstimateTotal(estimate *model.Estimate) float64 {
	return calcMiniTotals(estimate.Minis)
}

// QuoteTotal takes in a quote and calculates the total
func QuoteTotal(quote *model.Quote) float64 {
	return calcMiniTotals(quote.Minis)
}

func calcMiniTotals(minis []*model.MiniQuantity) float64 {
	total := 0.00

	// add up the cost for all minis
	for _, mini := range minis {
		qty := float64(mini.Quantity)
		cost := mini.Mini.Cost
		total = total + (qty * cost)

		// add up the options and multiply them for each mini
		for _, option := range mini.Options {
			total = total + (qty * option.Cost)
		}
	}

	return total
}
