package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DmitriyKomarovCoder/go_homework/unique"
)

func ReadDate(line *[]string) error {
	inputStream := os.Stdin
	var err error

	if fileName := flag.Arg(0); fileName != "" {
		inputStream, err = os.Open(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err = inputStream.Close(); err != nil {
				log.Fatal(err)
			}
		}()

	}

	scanner := bufio.NewScanner(inputStream)

	for scanner.Scan() {
		*line = append(*line, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return err
}

func writeDate(line []string) {
	outputStream := os.Stdout
	var err error

	if fileName := flag.Arg(1); fileName != "" {
		outputStream, err = os.Create(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err = outputStream.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	writer := bufio.NewWriter(outputStream)

	for _, str := range line {
		if _, err = writer.WriteString(str + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	if err = writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

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
	if err := ReadDate(&line); err != nil {
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

	writeDate(outString)
}
