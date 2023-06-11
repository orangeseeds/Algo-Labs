package knapsack

import "testing"

func TestKnapsackDP(t *testing.T) {
	problem := []Item{
		{1, 10},
		{2, 40},
		{1, 20},
		{3, 20},
		{5, 60},
	}

	var capacity float32 = 10

    result := KnapsackDP(int(capacity), problem)
    t.Log("result: ", result)
}
