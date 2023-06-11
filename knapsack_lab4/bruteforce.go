package knapsack

import (
	"math"
)

type Item struct {
	Weight float32
	Profit float32
}

func intToBitStream(num, size int) []float32 {
	var bitStream []float32

	for i := size; i >= 0; i-- {
		bit := (num >> i) & 1
		bitStream = append(bitStream, float32(bit))
	}
	return bitStream
}

func KnapsackBF(capacity float32, items []Item) []float32 {
	bitStreams := [][]float32{}
	combinations := math.Pow(2, float64(len(items)))
	for i := 0; i < int(combinations); i++ {
		bitStreams = append(bitStreams, intToBitStream(i, len(items)-1))
	}

	var (
		maxProfit float32
		config    []float32
	)

	for i := range bitStreams {
		var (
			sumProfit float32
			sumWeight float32
		)
		for j, bit := range bitStreams[i] {
			sumProfit += float32(bit) * items[j].Profit
			sumWeight += float32(bit) * items[j].Weight
		}

		if sumWeight <= capacity && sumProfit > maxProfit {
			maxProfit = sumProfit
			config = bitStreams[i]
		}
	}
	return config
}

func KnapsackFractBF(capacity float32, items []Item) ([]float32, float32) {
	bitStreams := [][]float32{}
	combinations := math.Pow(2, float64(len(items)))
	for i := 0; i < int(combinations); i++ {
		bitStreams = append(bitStreams, intToBitStream(i, len(items)-1))
	}

	var (
		maxProfit float32
		config    []float32
	)

	for i := range bitStreams {
		var (
			sumProfit float32
			sumWeight float32
		)
		for j, bit := range bitStreams[i] {
			remaining := capacity - sumWeight
			if items[j].Weight <= remaining {
				sumProfit += float32(bit) * items[j].Profit
				sumWeight += float32(bit) * items[j].Weight
			} else {
				sumProfit += (remaining / items[j].Weight) * items[j].Profit
				sumWeight += remaining

				bitStreams[i][j] = remaining / items[j].Weight
			}
		}

		if sumWeight <= capacity && sumProfit > maxProfit {
			maxProfit = sumProfit
			config = bitStreams[i]
		}
	}
	return config, maxProfit
}
