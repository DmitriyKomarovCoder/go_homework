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
	return str[opts.Strings:]
}

func uniqWorker(outString []string, lastStr string, count int, opts Options) []string {
	switch {
	case opts.Count:
		outString = append(outString, strconv.Itoa(count)+" "+lastStr)

	case opts.Double:
		if count > 1 {
			outString = append(outString, lastStr)
		}

	case opts.Unique:
		if count == 1 {
			outString = append(outString, lastStr)
		}

	default:
		outString = append(outString, lastStr)
	}
	return outString
}

func Unique(line []string, opts Options) ([]string, error) {
	var strPotential string
	outString := []string{}

	if opts.Count && opts.Double || opts.Double && opts.Unique || opts.Count && opts.Unique {
		return outString, errors.New("errors: No correct use flags")
	}

	count := 1
	for i := 1; i < len(line); i++ {
		lastStr := cut(line[i-1], opts)
		currentStr := cut(line[i], opts)

		if opts.Ignorant {
			currentStr = strings.ToUpper(currentStr)
			lastStr = strings.ToUpper(lastStr)
		}

		if currentStr == lastStr {
			if count == 1 {
				strPotential = line[i-1]
			}
			count++
			continue
		}

		outString = uniqWorker(outString, strPotential, count, opts)
		strPotential = line[i]
		count = 1
	}
	outString = uniqWorker(outString, strPotential, count, opts)
	return outString, nil
}
