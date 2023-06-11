package knapsack

import "testing"

func TestKnapsackGreedy(t *testing.T) {
	problem := []Item{
		{10, 60},
		{20, 100},
		{30, 120},
	}

	var capacity float32 = 50

	profit := KnapsackGreedy(capacity, problem)
	t.Logf("%v", profit)
}
