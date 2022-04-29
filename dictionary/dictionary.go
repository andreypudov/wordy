package dictionary

import (
	"github.com/magiconair/properties"
)

type Dictionary map[string]string

func Get(filename string) (Dictionary, error) {
	dictionary, err := properties.LoadFile(filename, properties.UTF8)
	return dictionary.Map(), err
}
