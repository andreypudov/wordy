package learning

import (
	"wordy/dictionary"
	"wordy/statistics"
)

type Teacher func(dictionary dictionary.Dictionary, statistics statistics.Statistics)

func Learn(dictionary dictionary.Dictionary, statistics statistics.Statistics) statistics.Statistics {
	var teacher Teacher

	teacher = TeacherCards
	teacher(dictionary, statistics)

	return statistics
}
