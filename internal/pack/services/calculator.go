package services

import (
	"repartners/internal/pack"
	"sort"
)

// Calculate - this code won't work with packs not divisible by smallest pack sizes. Like if the pack sizes were: 200, 300, 500.
// the function does calculate the number of packs needed so that no more items is sent, and the number of packs is as low as possible
func Calculate(items int, packs []pack.PackSize) map[int]int {
	result := make(map[int]int)

	sort.Slice(packs, func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	})

	smallestSize := packs[len(packs)-1].Size

	for _, packSize := range packs {
		curPacks := items / packSize.Size
		if curPacks > 0 {
			result[packSize.Size] = curPacks
		}
		items = items % packSize.Size

		if packSize.Size-items < smallestSize {
			result[packSize.Size]++
			items = 0
		}
	}

	return result
}
