package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/DmitriyKomarovCoder/go_homework/stream"
	"github.com/DmitriyKomarovCoder/go_homework/unique"
)

func main() {
	var opts unique.Options
	flag.BoolVar(&opts.Count, "c", false, unique.InfoCount)
	flag.BoolVar(&opts.Double, "d", false, unique.InfoDouble)
	flag.BoolVar(&opts.Unique, "u", false, unique.InfoUnique)
	flag.IntVar(&opts.Fields, "f", 0, unique.InfoField)
	flag.IntVar(&opts.Strings, "s", 0, unique.InfoString)
	flag.BoolVar(&opts.Ignorant, "i", false, unique.InfoIgnorant)
	flag.Parse()

	if len(flag.Args()) > 2 {
		fmt.Println(errors.New("error: too many arguments"))
		unique.InfoErrorsInput()
		return
	}

	line := []string{}
	if err := stream.ReadDate(&line); err != nil {
		fmt.Println(err)
		unique.InfoErrorsInput()
		return
	}

	outString, err := unique.Unique(line, opts)
	if err != nil {
		fmt.Println(err)
		unique.InfoErrorsInput()
		return
	}

	stream.WriteData(outString)
}
