package unique

import (
	"errors"
	"fmt"
	"strings"
)

const infoError = "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"

func InfoErrorsInput() {
	fmt.Println(infoError)
}

func cut(cutString string, opts Options) string {
	field := strings.Split(cutString, " ")

	if len(field) < opts.Fields {
		return ""
	}

	str := strings.Join(field[opts.Fields:], " ")

	if len(str) < opts.Strings {
		return ""
	}
	return str[opts.Strings:]

}

func Unique(line []string, opts Options) ([]string, error) {
	var cout int
	outString := []string{}
	if opts.Count && opts.Double || opts.Double && opts.Unique || opts.Count && opts.Unique {
		return outString, errors.New("errors: no correct use flags")
	}

	for i := 1; i < len(line); i++ {
		cutPrevString := cut(line[i-1], opts)
		cutString := cut(line[i], opts)

		if opts.Ignorant {
			cutPrevString = strings.ToUpper(cutPrevString)
			cutString = strings.ToUpper(cutString)
		}

		if cutPrevString == cutString {
			cout += 1
		}

		switch {
		case opts.Count:
		case opts.Double:
		case opts.Unique:
		default:
		}
	}
	return outString, nil
}
