package statistics

import (
	"os"
	"strconv"

	"github.com/magiconair/properties"
)

type Statistics map[string]int

func Get(filename string) (Statistics) {
	statistics := Statistics{}

	properties, err := properties.LoadFile(filename, properties.UTF8)
	if err != nil {
		return statistics
	}

	for word, statistic := range properties.Map() {
		value, err := strconv.Atoi(statistic)
		if err != nil {
			value = 0
		}

		statistics[word] = value
	}

	return statistics
}

func Update(statistics Statistics, filename string) (error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	for word, stat := range statistics {
		file.WriteString(word + "=" + strconv.Itoa(stat) + "\n")
	}

	return nil
}
