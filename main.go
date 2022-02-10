package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("example.txt")
	//fmt.Println(data)
	var tagRmv string
	var emptyString string
	var linkRmv string
	for i := 0; i < len(data); i++ {
		if data[i] != 10 {
			emptyString += string(data[i])
		}
		if data[i] == 10 {
			if strings.HasPrefix(emptyString, "#") {
				emptyString = ""
			} else {
				tagRmv += emptyString +"\n"
				emptyString=""
				
			}
		}
	}
	tagRmv+= emptyString +"\n"
	emptyString = ""
for i:=0; i<len(tagRmv); i++{
	if tagRmv[i] != 10{
		emptyString += string(tagRmv[i])
	}
	if tagRmv[i] == 10 {
		if strings.Contains(emptyString, "-") {
			emptyString = ""
		} else {
			linkRmv += emptyString +"\n"
			emptyString=""
			
		}
	}
}
fmt.Print(linkRmv)
	// scanner := bufio.NewScanner(data)
	// var a string

	// for scanner.Scan() {
	// 	//fmt.Println(scanner.Text())
	// 	a=scanner.Text()
	// 	a=strings.Trim(a, "#")
	// 	fmt.Println(a)
	// }
}
