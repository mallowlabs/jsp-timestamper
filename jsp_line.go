package main

import (
	"regexp"
	"strings"
)

type JSPLine struct {
	Line string
}

func NewJSPLine(line string) *JSPLine {
	return &JSPLine{Line: line}
}

func (self *JSPLine) Stamp(timestamp string) (string, bool) {
	if self.isScriptLine() {
		return self.stampScript(timestamp)
	} else if self.isStyleLine() {
		return self.stampStyle(timestamp)
	}
	return self.Line, false
}

func (self *JSPLine) isScriptLine() bool {
	return strings.Contains(self.Line, "<script") && strings.Contains(self.Line, "src=")
}

func (self *JSPLine) isStyleLine() bool {
	return strings.Contains(self.Line, "<link") && strings.Contains(self.Line, "href=")
}

func (self *JSPLine) stampScript(timestamp string) (string, bool) {
	return self.replaceFile("script", timestamp)
}

func (self *JSPLine) stampStyle(timestamp string) (string, bool) {
	return self.replaceFile("style", timestamp)
}

func (self *JSPLine) replaceFile(attr string, timestamp string) (string, bool) {
	file := self.exstractFile(attr)
	if file == "" {
		return self.Line, false
	}
	if strings.Contains(file, "?") {
		return self.Line, false
	}
	file = strings.Replace(self.Line, file, file+"?"+timestamp, 1)
	return file, true
}

func (self *JSPLine) exstractFile(attr string) string {
	startIndex := len("src=\"")
	pat := regexp.MustCompile(`src=\"(.+?)\"`)
	if attr == "style" {
		startIndex = len("href=\"")
		pat = regexp.MustCompile(`href=\"(.+?)\"`)
	}
	bytes := pat.Find([]byte(self.Line))
	if bytes == nil {
		return ""
	}
	matched := string(bytes)
	return matched[startIndex : len(matched)-1]
}
