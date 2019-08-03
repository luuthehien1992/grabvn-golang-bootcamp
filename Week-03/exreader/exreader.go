package exreader

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type ExReader struct {
	*bufio.Reader
}

var ConsoleReadError =  errors.New("Can't read console input")
var NumberInvalidError =  errors.New("Number invalid")

func (reader *ExReader) ReadInt32() (int32, error) {
	text, err := reader.ReadString('\n')

	if err != nil {
		return 0, ConsoleReadError
	}

	text = strings.TrimSuffix(text, "\n")
	i64, err := strconv.ParseInt(text, 10, 0)

	if err != nil {
		return 0, NumberInvalidError
	}

	return int32(i64), nil
}

func (reader *ExReader) ReadText() (string, error) {
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", ConsoleReadError
	}

	return strings.TrimSuffix(text, "\n"), nil
}

