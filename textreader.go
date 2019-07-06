package take5

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadLines(f func(string)) {
	reader, err := NewTextReader()
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for {
		line, isEof, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		if isEof {
			break
		}
		f(line)
	}
}

func OpenInputFile() (*os.File, error) {
	if len(os.Args) < 2 {
		return os.Stdin, nil
	}
	return os.Open(os.Args[1])
}

type TextReader struct {
	file   *os.File
	reader *bufio.Reader
}

func NewTextReader() (*TextReader, error) {
	var file *os.File
	var err error

	file, err = OpenInputFile()
	if err != nil {
		return nil, err
	}

	self := new(TextReader)
	self.file = file
	self.reader = bufio.NewReaderSize(file, 4096)
	return self, nil
}

func (self *TextReader) Close() {
	self.file.Close()
}

func (self *TextReader) ReadLine() (line string, isEof bool, err error) {
	isEof = false
	ispre := true
	slist := make([]string, 0, 1)
	for ispre {
		var buf []byte
		buf, ispre, err = self.reader.ReadLine()
		if err == io.EOF {
			isEof = true
			err = nil
			break
		} else if err != nil {
			line = ""
			isEof = false
			return
		}
		slist = append(slist, string(buf))
	}
	line = strings.Join(slist, "")
	return
}
