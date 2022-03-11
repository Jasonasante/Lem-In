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
	visited  int
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
}

// function to assign the start room for ants.
var (
	Start    *room
	lenStart int
)

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
	lenStart = len(Start.nextRoom)
	roomPaths = make([]string, lenStart)
}

var End *room

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
			End = ele
		}
	}
}

// find path

var count int // potPath     string
// ifRoomFound int

var roomPaths []string

// func allPaths(r *room) {
// 	room := r
// 	nextRoom := r.nextRoom
// 	if room.end {
// 		potPath += room.name
// 		allPaths(verifyPath(potPath))
// 	}
// 	for _, rooms := range nextRoom {
// 		if strings.Contains(potPath, rooms.name) {
// 			ifRoomFound++
// 		}
// 	}
// 	if ifRoomFound < len(nextRoom) {
// 		for j, rooms := range nextRoom {
// 			if !strings.Contains(potPath, rooms.name) {
// 				potPath += room.name + ","
// 				ifRoomFound = 0
// 				allPaths(nextRoom[j])
// 			} else {
// 				continue
// 			}
// 		}
// 	} else {
// 		potPath += room.name
// 		allPaths(verifyPath(potPath))
// 	}
// }

func allPaths(r *room) {
	prevRoom := r
	nextRoom := r.nextRoom
	fmt.Println(roomPaths)
	if prevRoom.end {
		roomPaths[count] += prevRoom.name
		Start.visited = 0
		verifyPath(roomPaths[count])
		count++
		allPaths(Start)
	}
	for i, rooms := range nextRoom {
		if count < lenStart {
			if !strings.Contains(roomPaths[count], rooms.name) && (rooms.visited == 0) {
				roomPaths[count] += prevRoom.name + ","
				prevRoom.visited = 1
				// fmt.Print(prevRoom.name + "/")
				// fmt.Println(prevRoom.visited)
				allPaths(nextRoom[i])
			}
		}
	}
}

// make function to make sure we dont repear path
// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false
// }
var (
	potPath    string
	returnPath string
)

func otherPaths(r *room) {
	returnPath := ""
	prevRoom := r
	nextRoom := r.nextRoom
	if prevRoom.end {
		returnPath += prevRoom.name
	}
	for i, rooms := range nextRoom {
		if !strings.Contains(potPath, rooms.name) && (rooms.visited == 0) {
			returnPath += prevRoom.name + ","
			prevRoom.visited = 1
			otherPaths(nextRoom[i])
		}
	}
}

func verifyPath(s string) {
	// fmt.Println(string(s[0]))
	// fmt.Println(string(s[(len(s)-1)]))
	var sent *room
	returnPath = strings.TrimRight(s, string(s[4]))
	fmt.Println("other part " + returnPath)
	//fmt.Println("potpath " + potPath)
	// fmt.Println("s string " +s)
	if (string(s[0]) == Start.name) && (string(s[(len(s)-1)]) == End.name) {

		for _, roomele := range roomList {
			if string(s[4]) == roomele.name {
				roomele.visited = 1
			}
			if string(s[2]) == roomele.name {
				sent = roomele
			}
		}
		for _, ele := range sent.nextRoom {
			if ele.visited == 0 {
				sent = ele
				// fmt.Println("other part " + returnPath)
				otherPaths(sent)
				// fmt.Println("potpath " + potPath)

			}
		}

	} else {
		roomPaths[count] = ""
	}

	// 	if contains(roomPaths, s) {
	// 		for _, roomele := range roomList {
	// 			if string(s[0]) == roomele.name {
	// 				for j, nextroomele := range roomele.nextRoom {
	// 					if strings.Contains(potPath, nextroomele.name) {
	// 						potPath = strings.TrimRight(s, nextroomele.name)
	// 						allPaths(roomele.nextRoom[j+1])
	// 					}
	// 				}
	// 			}
	// 		}
	// 		// for i := 0; i >= len(s)-1; i++ {
	// 		// 	if i%2 == 0 {
	// 		// 		for _, roomele := range roomList {
	// 		// 			if string(s[i]) == roomele.name {
	// 		// 				if roomele.end {

	// 		// 					// for _, nextroomele := range roomele.nextRoom {
	// 		// 					// 	if string(s[i+2]) != nextroomele.name && !strings.Contains(potPath, nextroomele.name) {
	// 		// 					// 		potPath = strings.TrimRight(s, string(s[i+2]))
	// 		// 					// 		allPaths(nextroomele)
	// 		// 					// 	} else {
	// 		// 					// 		continue
	// 		// 					// 	}
	// 		// 					// }
	// 		// 				}
	// 		// 			}
	// 		// 		}
	// 		// 	}
	// 		// }
	// 	} else if !contains(roomPaths, s) {
	// 		for _, pathrooms := range roomPaths {
	// 			for i := range s {
	// 				if i > 0 && i < len(s)-1 {
	// 					if strings.Contains(pathrooms, string(s[i+2])) {
	// 						for _, roomele := range roomList {
	// 							if string(s[i]) == roomele.name {
	// 								for _, nextroomele := range roomele.nextRoom {
	// 									if !strings.Contains(s, nextroomele.name) {
	// 										potPath = strings.TrimRight(s, string(s[i+2]))
	// 										allPaths(nextroomele)
	// 									}
	// 								}
	// 							}
	// 						}
	// 					} else {
	// 						roomPaths[count] = s
	// 						potPath = ""
	// 						count++
	// 						ifRoomFound = 0
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return Start
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
	linkRooms()
	assignStart()
	assignEnd()
	allPaths(Start)
}
