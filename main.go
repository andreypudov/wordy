package main

import (
	"log"
	"wordy/dictionary"
	"wordy/learning"
	"wordy/statistics"
)

func main() {
	learningstat := statistics.Get("statistics.txt")
	dictionary, err := dictionary.Get("dictionary.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	learning.Learn(dictionary, learningstat)

	err = statistics.Update(learningstat, "statistics.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
}
