package take5

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadWords(f func(string)) {
	buf := make([]rune, 0)
	grp := 'S'

	ReadRunes(func(c rune) {
		g := 'S'
		if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
			g = 'A'
		} else if '0' <= c && c <= '9' {
			g = 'N'
		} else if '~' < c {
			g = 'O'
		}

		if g != grp && 0 < len(buf) {
			f(string(buf))
			buf = make([]rune, 0)
		}

		grp = g

		if g == 'S' {
			f(string(c))
		} else {
			buf = append(buf, c)
		}
	})

	if 0 < len(buf) {
		f(string(buf))
	}
}

func ReadRunes(f func(rune)) {
	reader, err := NewTextReader()
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	for {
		c, isEof, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		if isEof {
			break
		}
		f(c)
	}
}

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
	self.reader = bufio.NewReader(file)
	return self, nil
}

func (self *TextReader) Close() {
	self.file.Close()
}

func (self *TextReader) ReadRune() (char rune, isEof bool, err error) {
	isEof = false
	char, _, err = self.reader.ReadRune()
	if err == io.EOF {
		isEof = true
		err = nil
	}
	return
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
