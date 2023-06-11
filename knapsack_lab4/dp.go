package knapsack

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func KnapsackDP(capacity int, items []Item) []int {
	table := make([][]int, len(items)+1)
	for i := range table {
		table[i] = make([]int, capacity+1)
	}

	for i := 0; i <= len(items); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if int(items[i-1].Weight) > j {
				table[i][j] = table[i-1][j]
			} else {
				if int(items[i-1].Profit)+table[i-1][j-int(items[i-1].Weight)] > table[i-1][j] {
					table[i][j] = int(items[i-1].Profit) + table[i-1][j-int(items[i-1].Weight)]
				} else {
					table[i][j] = table[i-1][j]
				}
			}
		}
	}
    result := solve(table, capacity+1, len(items)+1, items)
	// displayTable(table)
	return result
}

func solve(table [][]int, w, h int, items []Item) []int {
	taken := []int{}
	i, j := w-1, h-1
	profit := table[j][i]
	for i > 0 && j > 0 {
		if table[j][i-1] == table[j][i] {
			i -= 1
		} else if table[j-1][i] == table[j][i] {
			j -= 1
		} else {
			taken = append(taken, j-1)
            profit = profit - int(items[j-1].Profit)
			j -= 1
            index := slices.IndexFunc(table[j], func(item int) bool {
				return item == profit
			})
            if index == -1 {
                j -= 1
            } else {
                i = index
            }
		}
	}
    return taken
}

func displayTable(table [][]int) {
	for i := range table {
		for j := range table[i] {
			fmt.Print(table[i][j], "\t")
		}
		fmt.Print("\n")
	}
}
