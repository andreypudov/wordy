package learning

import (
	"bufio"
	"fmt"
	"os"
	"wordy/dictionary"
	"wordy/statistics"
	"wordy/utils"
)

func TeacherCards(dictionary dictionary.Dictionary, statistics statistics.Statistics) {
	sorted_statistics := utils.SortStatisticsByDictionary(statistics, dictionary)

	// infinit learning loop
	for {
	for _, stat := range sorted_statistics {
		translation := dictionary[stat.Key]
		statistics[stat.Key] += 1

		fmt.Println("Word: \t\t", stat.Key)
		fmt.Print("Press any key to continue or Escape to exit...")

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case '\u001b':
			clear_previous_line()
			fmt.Println("Translation: \t", translation)
			fmt.Println("Bye")
			return
		}

		clear_previous_line()
		fmt.Println("Translation: \t", translation)
		fmt.Println("")
	}
	}
}

func clear_previous_line() {
	fmt.Print("\033[A")
	fmt.Print("\033[G\033[K")
}
