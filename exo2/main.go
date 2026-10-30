package main

import (
	"fmt"
	"os"
	"sort"
)

type tagCount struct {
	tag   string
	count int
}

func main() {
	tags := os.Args[1:]

	if len(tags) == 0 {
		fmt.Println("Aucun tag fourni")
		return
	}

	counts := make(map[string]int)

	for _, tag := range tags {
		counts[tag]++
	}

	var results []tagCount

	for tag, count := range counts {
		if count > 1 {
			results = append(results, tagCount{
				tag:   tag,
				count: count,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})

	fmt.Println("Tags présents plusieurs fois :")

	for _, result := range results {
		fmt.Printf("%s : %d\n", result.tag, result.count)
	}
}
