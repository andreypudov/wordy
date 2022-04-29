package learning

import (
	"fmt"
	"wordy/dictionary"
	"wordy/statistics"
)

func Learn(dictionary dictionary.Dictionary, statistics statistics.Statistics) statistics.Statistics {
	for word, translation := range dictionary {
		stat := statistics[word]
		fmt.Println(word, " := ", translation, " [", stat, "]")

		statistics[word] += 1
	}

	return statistics
}
