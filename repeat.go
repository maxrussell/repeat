package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const optimalSize = 1024 * 4

var (
	bufferSize = flag.Int("buffer", optimalSize, "the minimum size in bytes of the repeated output")
	byteMode   = flag.Bool("byte", false, "if set, the program will interpret the args as a sequence of numerical bytes to repeat")
)

func main() {
	flag.Parse()

	var message string
	if byteMode != nil && *byteMode {
		args := flag.Args()
		var err error
		message, err = parseByteArgs(args)
		if err != nil {
			panic(err)
		}
	} else {
		message = strings.Join(flag.Args(), " ")
	}

	if len(message) < *bufferSize {
		sizeRatio := float64(*bufferSize) / float64(len(message))
		message = strings.Repeat(message, int(sizeRatio))
	}

	for {
		fmt.Print(message)
	}
}

func parseByteArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("byte mode requires at least one integer parameter from 0-255")
	}

	buf := bytes.Buffer{}
	for _, arg := range args {
		parsed, err := strconv.Atoi(arg)
		if err != nil {
			return "", fmt.Errorf("failed to parse `%s` as an integer", arg)
		} else if parsed >= 256 {
			return "", errors.New("byte argument must be between 0 and 255 (inclusive)")
		}

		buf.WriteByte(byte(parsed))
	}

	return buf.String(), nil
}
