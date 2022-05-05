package utils

import (
	"sort"
	"wordy/dictionary"
	"wordy/statistics"
)

type Statistics struct {
	Key   string
    Value int
}

type StatisticsList []Statistics

func (p StatisticsList) Len() int { return len(p) }
func (p StatisticsList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p StatisticsList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func SortStatistics(original_statistics statistics.Statistics) StatisticsList {
	sorted_pairs := make(StatisticsList, len(original_statistics))
	index := 0

	for key, value := range original_statistics {
		sorted_pairs[index] = Statistics{key, value}
		index++
	}

	sort.Sort(sorted_pairs)

	return sorted_pairs
}

func SortStatisticsByDictionary(original_statistics statistics.Statistics, dictionary dictionary.Dictionary) StatisticsList {
	filled_statistics := statistics.Statistics{}

	for word, translation := range original_statistics {
		filled_statistics[word] = translation
	}

	for word, _ := range dictionary {
		if _, exists := filled_statistics[word]; !exists {
			filled_statistics[word] = 0
		}
	}

	return SortStatistics(filled_statistics)
}
