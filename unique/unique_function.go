package unique

import "fmt"

const infoError = "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

func InfoErrorsInput() {
	fmt.Println(infoError)
	return
}

func Unique(line *[]string, opts Options) error {
	for i := 0; i < len(*line); i++ {

	}
	return nil
}
