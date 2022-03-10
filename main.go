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
	start    bool
	end      bool
}

var roomList []*room

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
			roomList = append(roomList, rooms)
			emptyString = ""
		}
	}
}

// function to assign the start room for ants.
var Start *room

func assignStart() {
	data, _ := os.Open("example.txt")
	var getStart []string
	var startLine string
	startInfo := bufio.NewScanner(data)
	for startInfo.Scan() {
		getStart = append(getStart, startInfo.Text())
	}

	for i := range getStart {
		if getStart[i] == "##start" {
			startLine = getStart[i+1]
		}
	}
	a := strings.Split(startLine, " ")
	for _, ele := range roomList {
		if ele.name == a[0] {
			ele.start = true
		}
	}
	for i := range roomList {
		if roomList[i].start {
			Start = roomList[i]
		}
	}
}

// function to assign the end room for ants
func assignEnd() {
	data, _ := os.Open("example.txt")
	var getEnd []string
	var endLine string
	// this is to get coords by removing # and -
	endInfo := bufio.NewScanner(data)
	for endInfo.Scan() {
		getEnd = append(getEnd, endInfo.Text())
	}

	for i := range getEnd {
		if getEnd[i] == "##end" {
			endLine = getEnd[i+1]
		}
	}
	a := strings.Split(endLine, " ")
	for _, ele := range roomList {
		if ele.name == a[0] {
			ele.end = true
			ele.nextRoom = nil
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
				for k := range roomList {
					for o := range roomList {
						if string(links[i][j-1]) == roomList[k].name && roomList[o].name == string(links[i][j+1]) {
							roomList[k].nextRoom = append(roomList[k].nextRoom, roomList[o])
						} else if string(links[i][j+1]) == roomList[k].name && roomList[o].name == string(links[i][j-1]) {
							roomList[k].nextRoom = append(roomList[k].nextRoom, roomList[o])
						}
					}
				}
			}
		}
	}
	// for _, ele := range roomList {
	// 	fmt.Println("roomList of rooms", *ele)
	// }
}

// find path

var (
	roomPaths []string = make([]string, 1)
	count     int
	potPath   string
)

func allPaths(r *room) {
	room := r
	nextRoom := r.nextRoom
	// fmt.Println(roomPaths)
	if room.end {
		potPath += room.name
		verifyPath(potPath)

	} else {
		for j, rooms := range nextRoom {
			if !strings.Contains(potPath, rooms.name) {
				potPath += room.name + ","
				room = nextRoom[j]
				allPaths(room)
			} else if !strings.Contains(potPath, rooms.name) && len(nextRoom) == 1 {
				potPath += room.name
				potPath += rooms.name
				verifyPath(potPath)
			} else {
				potPath += room.name + ","
				room = nextRoom[j+1]
				allPaths(room)
			}
		}
	}
}

// make function to make sure we dont repear path
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func verifyPath(s string) {
	if contains(roomPaths, s) {
		for i := len(s) - 1; i == 0; i-- {
			for _, roomele := range roomList {
				if string(s[i]) == roomele.name {
					if roomele.nextRoom != nil {
						for _, nextroomele := range roomele.nextRoom {
							if nextroomele.name != string(s[i+1]) {
								potPath = strings.TrimRight(potPath, string(potPath[i+1]))
								allPaths(nextroomele)
							}
						}
					}
				}
			}
		}
	} else {
		roomPaths[count] = potPath
		potPath = ""
		count++
		allPaths(Start)
	}
	fmt.Println(potPath)
}

// find shortest route

// place rooms in grid
func grid() {
	var grid [30][30]string
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid); column++ {
			grid[row][column] = " "
		}
	}
	for i := range roomList {
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid); column++ {
				if row == roomList[i].row-1 && column == roomList[i].column-1 {
					grid[row][column] = "[" + roomList[i].name + "]"
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
	allPaths(Start)
}
