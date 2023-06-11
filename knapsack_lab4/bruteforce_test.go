package knapsack

import (
	"testing"
)

func TestKnapsackBF(t *testing.T) {

    var profit float32
	problem := []Item{
		{7, 42},
		{3, 12},
		{4, 40},
		{5, 25},
	}

	var capacity float32 = 10

	config := KnapsackBF(capacity, problem)
	t.Logf("%v", config)
    for i  := range config {
        profit += config[i] * problem[i].Profit
    }
	t.Logf("%v", profit)
}

func TestKnapsackFractBF(t *testing.T){
	problem := []Item{
        {10, 60},
        {20, 100},
        {30, 120},
	}

	var capacity float32 = 50

	config, profit := KnapsackFractBF(capacity, problem)
	t.Logf("%v", config)
	t.Logf("%v", profit)
}
