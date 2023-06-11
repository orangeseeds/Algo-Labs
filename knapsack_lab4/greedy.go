package knapsack

import (
	"sort"
)

func KnapsackGreedy(capacity float32, items []Item) float32 {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Profit/items[i].Weight > items[j].Profit/items[j].Weight
	})

	var (
		sumWeight float32
		sumProfit float32
	)

	for i := 0; i < len(items); i++ {
        remaining := capacity - sumWeight
		if items[i].Weight <= remaining {
			sumWeight += items[i].Weight
			sumProfit += items[i].Profit
		} else {
			sumProfit += items[i].Profit * (remaining / items[i].Weight)
			break
		}
	}
    return sumProfit
}
