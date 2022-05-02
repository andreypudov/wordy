package learning

import (
	"fmt"
	"wordy/dictionary"
	"wordy/statistics"
	"wordy/utils"
)

// 1. Sort words by rating

func Learn(dictionary dictionary.Dictionary, statistics statistics.Statistics) statistics.Statistics {
	sorted_statistics := utils.SortStatistics(statistics)

	for _, stat := range sorted_statistics {
		translation := dictionary[stat.Key]
		fmt.Println(stat.Key, " := ", translation, " [", stat.Value, "]")

		statistics[stat.Key] += 1
	}

	return statistics
}
