package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {

	//err := os.Setenv("TITLE", "feat(dot/rpc): implement chain_subscribeAllHeads RPC")
	//if err != nil {
	//	return
	//}
	//fmt.Println("TITLE:", os.Getenv("TITLE"))
	var match, _ = regexp.MatchString(".+\\(.+\\)\\:.+", os.Getenv("TITLE"))
	//fmt.Println(match)
	if match == false{
		os.Exit(1)
	}


	//data, err := ioutil.ReadFile(".github/workflows/BODY_TEMPLATE.md")
	//if err != nil {
	//	fmt.Println("File reading error", err)
	//	os.Exit(1)
	//	return
	//}
	//err = os.Setenv("BODY", string(data))
	//	fmt.Println(os.Getenv("BODY"))
	data := os.Getenv("BODY")
	data1 := strings.Split(data, "")
	//data1 := strings.Split(string(data), "")
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
	//fmt.Println("Joined line: ", joinedLine)
	//var validPR= regexp.MustCompile(`## Changes.*- .+[A-Za-z0-9].+## Tests.*- .+[A-Za-z0-9].+## Issues.*- .+[A-Za-z0-9].+## Primary Reviewer.*- .+[A-Za-z0-9].+`)
	////fmt.Println(validPR.MatchString(joinedLine))
	//validPR.MatchString(joinedLine)
	var validPR,_= regexp.MatchString(`## Changes.*-.+[A-Za-z0-9].+## Tests.*-.+[A-Za-z0-9].+## Issues.*-.+[A-Za-z0-9].+## Primary Reviewer.*-.+[A-Za-z0-9].+`,joinedLine)
	//fmt.Println(validPR.MatchString(joinedLine))
	if validPR == false{
		os.Exit(1)
	}
}