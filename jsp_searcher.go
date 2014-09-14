package main

import (
	"os"
	"path/filepath"
	"strings"
)

type JSPSearcher struct {
	Path string
}

func NewJSPSearcher(path string) *JSPSearcher {
	return &JSPSearcher{Path: path}
}

func (self *JSPSearcher) List() ([]*JSPFile, error) {
	if _, err := os.Stat(self.Path); os.IsNotExist(err) {
		return nil, err
	}

	jspFiles := []*JSPFile{}
	err := filepath.Walk(self.Path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if strings.ToLower(filepath.Ext(path)) != ".jsp" {
				return nil
			}

			jspFiles = append(jspFiles, &JSPFile{Path: path})

			return nil
		})
	return jspFiles, err
}
