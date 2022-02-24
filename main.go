package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type room struct {
	name     string
	column   int
	row      int
	nextRoom []*room
	prevRoom []*room
	start    bool
	end      bool
	visited  int
}

var list []*room


// to initialise rooms with their own address
func getRooms() {
	data, _ := os.Open("example.txt")
	// fmt.Println(data)
	var emptyString string
	var getCoOrd string
	// this is to get coords by removing # and -
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		// fmt.Println(scanner.Text())
		if strings.Contains(getCoOrds.Text(), "#") {
			emptyString = ""
		} else if strings.Contains(getCoOrds.Text(), "-") {
			emptyString = ""
		} else {
			emptyString = getCoOrds.Text() + "\n"
			getCoOrd += emptyString
			emptyString = ""
		}
	}
	// fmt.Print(linkRmv)
	var a []string
	var rooms *room
	var rowInt int
	var columnInt int

	// this is to add co-ordinates to their respective room struct
	for i := 0; i < len(getCoOrd); i++ {
		if getCoOrd[i] != 10 {
			emptyString += string(getCoOrd[i])
		}
		if getCoOrd[i] == 10 {
			a = strings.Split(emptyString, " ")
			columnInt, _ = strconv.Atoi(a[1])
			rowInt, _ = strconv.Atoi(a[2])
			rooms = &room{name: a[0]}
			rooms.column = columnInt
			rooms.row = rowInt
			list = append(list, rooms)
			emptyString = ""
		}
	}
}

//function to assign the start room for ants.
func assignStart() {
	data, _ := os.Open("example.txt")
	var getStart []string
	var startLine string
	startInfo := bufio.NewScanner(data)
	for startInfo.Scan() {
		getStart = append(getStart, startInfo.Text())
	}

	for i :=range getStart{
		if getStart[i]=="##start"{
			startLine=getStart[i+1]
		}
	}
	a:= strings.Split(startLine," ")
	for _,ele:=range list{
		if ele.name==a[0]{
			ele.start=true
		}
	}
}

//function to assign the end room for ants
func assignEnd() {
	data, _ := os.Open("example.txt")
	var getEnd []string
	var endLine string
	// this is to get coords by removing # and -
	endInfo := bufio.NewScanner(data)
	for endInfo.Scan() {
		getEnd = append(getEnd, endInfo.Text())
	}

	for i :=range getEnd{
		if getEnd[i]=="##end"{
			endLine=getEnd[i+1]
		}
	}
	a:= strings.Split(endLine," ")
	for _,ele:=range list{
		if ele.name==a[0]{
			ele.end=true
		}
	}
}

// this links the room to their respective next room(s)
func linkRooms() {
	data, _ := os.Open("example.txt")
	// fmt.Println(data)
	var emptyString string
	var links []string
	// this is to get coords by removing # and -
	linksInfo := bufio.NewScanner(data)
	for linksInfo.Scan() {
		// fmt.Println(scanner.Text())
		if strings.Contains(linksInfo.Text(), "#") {
			emptyString = ""
		} else if strings.Contains(linksInfo.Text(), " ") {
			emptyString = ""
		} else {
			emptyString = linksInfo.Text()
			links = append(links, emptyString)
			emptyString = ""
		}
	}
	for i := range links {
		for j := range links[i] {
			if links[i][j] == '-' {
				for k := range list {
					for o := range list {
						if string(links[i][j-1]) == list[k].name && list[o].name == string(links[i][j+1]) {
							length := len(list[k].nextRoom)
							list[k].nextRoom = append(list[k].nextRoom, list[o])
							list[k].nextRoom[length].prevRoom = append(list[k].nextRoom[length].prevRoom, list[k])
						} else if string(links[i][j+1]) == list[k].name && list[o].name == string(links[i][j-1]) {
							length := len(list[k].nextRoom)
							list[k].nextRoom = append(list[k].nextRoom, list[o])
							list[k].nextRoom[length].prevRoom = append(list[k].nextRoom[length].prevRoom, list[k])
						}
					}
				}
			}
		}
	}
	// for _, ele := range list {
	// 	fmt.Println("list of rooms", *ele)
	// 	for _, room := range ele.prevRoom {
	// 		fmt.Println(*room)
	// 	}
	// }
}

// find path


// find shortest route

// place rooms in grid
func grid() {
	var grid [30][30]string
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			grid[row][column] = " "
		}
	}
	for i := range list {
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid); column++ {
				if row == list[i].row-1 && column == list[i].column-1 {
					grid[row][column] = "[" + list[i].name + "]"
				}
			}
		}
	}
	for i := range grid {
		for _, ele := range grid[i] {
			fmt.Print(ele)
		}
		fmt.Println()
	}
}

func main() {
	getRooms()
	assignStart()
	assignEnd()
	linkRooms()
}

// find path

// find shortest route
