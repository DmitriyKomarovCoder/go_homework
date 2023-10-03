package stream

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func ReadDate() ([]string, error) {
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
	line := []string{}
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line, err
}

func WriteData(line []string) {
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
