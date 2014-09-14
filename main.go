package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: jsp-timestamper <JSP_DIRECTORY_PATH>")
		os.Exit(1)
	}

	timestamp := GetCurrentTimestamp()
	jspSearcher := NewJSPSearcher(os.Args[1])
	jspFiles, err := jspSearcher.List()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	for _, f := range jspFiles {
		fmt.Println("Processing: " + f.Path)
		replaced, _ := f.Stamp(timestamp)
		if replaced > 0 {
			fmt.Println(strconv.Itoa(replaced) + " lines stamped.")
		}
	}
}
