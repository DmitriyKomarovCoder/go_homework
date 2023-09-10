package unique

import (
	"errors"
	"fmt"
	"strconv"
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

	return str
}

func Unique(line []string, opts Options) ([]string, error) {
	var cout int
	var isEqual bool
	outString := []string{}
	line = append(line, " ")
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
			cout++
			isEqual = true
		} else {
			isEqual = false
		}

		switch {
		case opts.Count:
			if !isEqual && cout > 0 {
				outString = append(outString, strconv.Itoa(cout+1)+" "+line[i-1])
				cout = 0
			} else if !isEqual && cout == 0 {
				outString = append(outString, strconv.Itoa(cout+1)+" "+line[i-1])
			}
		case opts.Double:
			if !isEqual && cout > 0 { // TO DO
				outString = append(outString, line[i-1])
				cout = 0
			}
		case opts.Unique:
			if !isEqual && cout == 0 {
				outString = append(outString, line[i-1])
			} else if !isEqual && cout > 0 {
				cout = 0
			}
		default:
			if !isEqual && cout > 0 {
				outString = append(outString, line[i-1])
			} else if !isEqual && cout == 0 {
				outString = append(outString, line[i-1])
			}
		}
	}
	return outString, nil
}
