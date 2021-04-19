package utils

import "github.com/myminicommission/api/graph/model"

// PriceForSize returns the float64 value of the associated incoming price
func PriceForSize(prices *model.Prices, size model.MiniSize) float64 {
	switch size {
	case model.MiniSizeTiny:
		return prices.Tiny
	case model.MiniSizeSmall:
		return prices.Small
	case model.MiniSizeMedium:
		return prices.Medium
	case model.MiniSizeLarge:
		return prices.Large
	case model.MiniSizeExtralarge:
		return prices.Extralarge
	case model.MiniSizeTitanic:
		return prices.Titanic
	}

	return 0.00 // in case we find nothing, set the value to zero
}
