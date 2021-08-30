package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {
	var match, _ = regexp.MatchString(".+\\(.+\\)\\:.+", os.Getenv("TITLE"))
	if match == false{
		os.Exit(1)
	}

	data := os.Getenv("BODY")
	data1 := strings.Split(data, "")
	data2 := make([]string, len(data1))
	for i := 0; i < len(data1); i++ {
		if data1[i] == "<" && data1[i+1] == "!" && data1[i+2] == "-" && data1[i+3] == "-" {
			for j := i + 4; j < len(data1); j++ {
				if data1[j] == "-" && data1[j+1] == "-" && data1[j+2] == ">" {
					i = j + 3
					break
				}
			}
		} else {
			data2[i] = data1[i]
		}
	}
	var result string
	for i := 0; i < len(data2); i++ {
		result = result + data2[i]
	}
	lineSplit := strings.Split(result, "\n")
	joinedLine := strings.Join(lineSplit, "")
	var validPR,_= regexp.MatchString(`## Changes.*-.+[A-Za-z0-9].+## Tests.*-.+[A-Za-z0-9].+## Issues.*-.+[A-Za-z0-9].+## Primary Reviewer.*-.+[A-Za-z0-9].+`,joinedLine)
	if validPR == false{
		os.Exit(1)
	}
}