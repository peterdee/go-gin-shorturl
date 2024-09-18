package utilities

import "github.com/nrednav/cuid2"

func GenerateCUID(length ...int) (string, error) {
	symbolsLength := 10
	if len(length) > 0 && length[0] > 0 {
		symbolsLength = length[0]
	}
	generator, generatorError := cuid2.Init(cuid2.WithLength(symbolsLength))
	if generatorError != nil {
		return "", generatorError
	}
	return generator(), nil
}
