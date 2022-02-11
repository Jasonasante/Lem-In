package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type startEnd struct {
	start *room 
	length int
}

type room struct {
	name     string
	nextRoom *room
}


//link rooms
func (s *startEnd) linkRoom(r *room) {
second:=s.start
s.start= r
s.start.nextRoom=second
s.length++

// var emptyString string
// var splitLink []string

// r = make([]room, len(s))

// 	for i := 0; i < len(s); i++ {		
// 		if s[i] == 10 {
// 			splitLink = strings.Split(emptyString,"-")
// 			emptyString = ""
// 			rooms[i].name = string(splitLink[0])
// 			rooms[i].nextRoom.name= splitLink[1]
// 		}
// 		if s[i] != 10 {
// 			emptyString += string(s[i])
// 		}
// 	}
// fmt.Print(rooms)	
}

func main() {
	data, _ := os.Open("example.txt")
	// fmt.Println(data)
	var emptyString string
	var links string
	//this is to get coords by removing # and -
	getLinks := bufio.NewScanner(data)
	for getLinks.Scan() {
		// fmt.Println(scanner.Text())
		if strings.Contains(getLinks.Text(), "#") {
			emptyString = ""
		} else if strings.Contains(getLinks.Text(), " ") {
			emptyString = ""
		} else {
			emptyString = getLinks.Text() + "\n"
			links += emptyString
			emptyString = ""

		}
	}
	roomsLink := room{}
	roomsLink.linkRoom(links) 

}
//find path 

//find shortest route










func visualiser() {
	var grid [30][30]string

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			grid[row][column] = " "
		}
	}
	data, _ := os.Open("example.txt")
	// fmt.Println(data)
	var emptyString string
	var linkRmv string
	//this is to get coords by removing # and -
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		// fmt.Println(scanner.Text())
		if strings.Contains(getCoOrds.Text(), "#") {
			emptyString = ""
		} else if strings.Contains(getCoOrds.Text(), "-") {
			emptyString = ""
		} else {
			emptyString = getCoOrds.Text() + "\n"
			linkRmv += emptyString
			emptyString = ""

		}
	}
	// fmt.Print(linkRmv)
	var a []string
	var rooms room
	var rowInt int
	var columnInt int
	//this is to place to rooms on grid
	for i := 0; i < len(linkRmv); i++ {
		if linkRmv[i] != 10 {
			emptyString += string(linkRmv[i])
		}
		if linkRmv[i] == 10 {
			a = strings.Split(emptyString, " ")
			columnInt, _ = strconv.Atoi(a[1])
			rowInt, _ = strconv.Atoi(a[2])
			rooms.name = a[0]
			for row := 0; row < len(grid); row++ {
				for column := 0; column < len(grid); column++ {
					if row == rowInt-1 && column == columnInt-1 {
						grid[row][column] =  "[" +rooms.name + "]"
					}
				}
			}
			emptyString = ""
		}
	}
	//this is to print grid
	for i := range grid {
		for _, ele := range grid[i] {
			fmt.Print(ele)
		}
		fmt.Println()
	}

}
