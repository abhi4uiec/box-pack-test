package service

import (
	"sort"
)

func CalculatePackSizes(boxSizes []int, numOfItems int) map[int]int {

	modTmp := make([]int, len(boxSizes))
	divTmp := make([]int, len(boxSizes))
	result := make(map[int]int)

	minOfModAndDiv, box := 0, 0
	quantity := 1

	sort.Ints(boxSizes)

	for quantity > 0 && numOfItems > 0 {
		for i := 0; i < len(boxSizes); i++ {
			divTmp[i] = numOfItems / boxSizes[i]
			modTmp[i] = numOfItems % boxSizes[i]

			if modTmp[i] == numOfItems && i == 1 && boxSizes[i-1] < numOfItems {
				box = boxSizes[i]
				break
			}

			if i == 0 || (divTmp[i]+modTmp[i]) < minOfModAndDiv {
				minOfModAndDiv = divTmp[i] + modTmp[i]
				box = boxSizes[i]
				quantity = divTmp[i]
			}
		}

		numOfItems = numOfItems - box

		if val, ok := result[box]; ok {
			result[box] = val + 1
		} else {
			result[box] = 1
		}
	}
	return result
}
