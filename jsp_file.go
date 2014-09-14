package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type JSPFile struct {
	Path string
}

func NewJSPFile(path string) *JSPFile {
	return &JSPFile{Path: path}
}

func (self *JSPFile) Stamp(timestamp string) (int, error) {
	lines, err := self.readLines()
	if err != nil {
		return 0, err
	}
	lines, replaced := self.stampLines(lines, timestamp)
	self.writeLines(lines)
	return replaced, nil
}

func (self *JSPFile) readLines() ([]string, error) {
	lines := []string{}

	fp, err := os.Open(self.Path)
	if err != nil {
		return lines, err
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, err := reader.ReadString('\n')
		lines = append(lines, line)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return lines, err
		}
	}
	return lines, err
}

func (self *JSPFile) stampLines(lines []string, timestamp string) ([]string, int) {
	var i int
	replaced := []string{}
	for _, l := range lines {
		newLine, isReplaced := self.stampLine(l, timestamp)
		if isReplaced {
			i++
		}
		replaced = append(replaced, newLine)
	}
	return replaced, i
}

func (self *JSPFile) stampLine(line string, timestamp string) (string, bool) {
	jspLine := NewJSPLine(line)
	return jspLine.Stamp(timestamp)
}

func (self *JSPFile) writeLines(lines []string) error {
	bytes := []byte(strings.Join(lines, ""))
	err := ioutil.WriteFile(self.Path, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
