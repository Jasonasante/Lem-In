package main

import (
	"bufio"
	"fmt"
	"log"
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

type Ant struct {
	name string
	path string
}

func getAnts() {
	data, _ := os.Open("example.txt")
	getants := bufio.NewScanner(data)
	line := 0

	for getants.Scan() {
		line++
		if line == 1 {
			a, err := strconv.Atoi(getants.Text())
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				log.Fatal()
			}
			fmt.Println(a)
		}
	}
}

var roomList []*room

// to initialise rooms with their own address
func getRooms() {
	data, _ := os.Open("example.txt")
	// fmt.Println(data)
	var emptyString string
	var getCoOrd string
	line:=0
	// this is to get coords by removing # and -
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		line++
		if line>1{
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
	}
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
	line:=0
	// this is to get coords by removing # and -
	linksInfo := bufio.NewScanner(data)
	for linksInfo.Scan() {
		line++
		if line>1{
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
	}

	for i := range links {
		for j := range links[i] {
			if links[i][j] == '-' {
				linkString := strings.Split(links[i], "-")
				// fmt.Println(linkString)
				for k := range roomList {
					for o := range roomList {
						if linkString[0] == roomList[k].name && roomList[o].name == linkString[1] {
							roomList[k].nextRoom = append(roomList[k].nextRoom, roomList[o])
						} else if linkString[1] == roomList[k].name && roomList[o].name == linkString[0] {
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
	roomPaths = make([]string, 5)
}


// function to assign the end room for ants
var End *room

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
			End = ele
		}
	}
}

// find path from start to end
var (
	count     int
	roomPaths []string
)

func allPaths(r *room) {
	prevRoom := r
	nextRoom := r.nextRoom
	visitedCounter := 0
	lenCounter := 0
	if prevRoom.end {
		roomPaths[count] += prevRoom.name
		roomPaths[count] = startEnd(roomPaths[count])
		Start.visited = 0
		End.visited = 0
		count++
		allPaths(Start)
	}
	for _, ele := range nextRoom {
		if ele.visited == 1 {
			visitedCounter++
		}
	}

	if visitedCounter == len(nextRoom) {
		prevRoom.visited = 1
		dEndNameSlice := strings.Split(roomPaths[count], ",")
		for _, room := range roomList {
			if len(dEndNameSlice) >= 2 {
				if dEndNameSlice[len(dEndNameSlice)-2] == room.name {
					dEndNameSlice = remove(dEndNameSlice, len(dEndNameSlice)-2)
					roomPaths[count] = strings.Join(dEndNameSlice, ",")
				}
			}
		}
	} else {
		for _, roomele := range nextRoom {
			lenCounter++
			if prevRoom == Start {
				for i, rooms := range nextRoom {
					if count < lenStart {
						rNamesSlice := strings.Split(roomPaths[count], ",")
						if !contains(rNamesSlice, rooms.name) && (rooms.visited == 0) && !strings.HasPrefix(rooms.name, "G") {
							roomPaths[count] += prevRoom.name + ","
							prevRoom.visited = 1
							allPaths(nextRoom[i])
						}
					}
				}
			}
			if roomele.end {
				if count < lenStart {
					roomPaths[count] += prevRoom.name + ","
					prevRoom.visited = 1
					allPaths(roomele)
				}
			} else if lenCounter == len(nextRoom) {
				for _, check := range nextRoom {
					for _, endNextRooms := range End.nextRoom {
						if check.name == endNextRooms.name && check.visited == 0 {
							if count < lenStart {
								rNamesSlice := strings.Split(roomPaths[count], ",")
								if !contains(rNamesSlice, check.name) && !strings.HasPrefix(check.name, "G") {
									if prevRoom != End {
										check.visited = 1
										endNextRooms.visited = 1
										prevRoom.visited = 1
										roomPaths[count] += prevRoom.name + ","
										roomPaths[count] += check.name + ","
										allPaths(End)
									}
								}
							}
						}
					}
				}
				for i, rooms := range nextRoom {
					if count < lenStart {
						rNamesSlice := strings.Split(roomPaths[count], ",")
						if !contains(rNamesSlice, rooms.name) && (rooms.visited == 0) && !strings.HasPrefix(rooms.name, "G") {
							if prevRoom != End {
								roomPaths[count] += prevRoom.name + ","
								prevRoom.visited = 1
								allPaths(nextRoom[i])
							}
						}
					}
				}
			}
		}
	}
}

// confirms whether a slice contains a certain value
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// removes a slice of string from an array at the given index
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// removes an invalid path
func startEnd(s string) string {
	a := strings.Split(s, ",")
	if a[0] == Start.name && a[len(a)-1] == End.name {
	} else {
		s = ""
	}
	return s
}

// removes empty paths from path's array
var finalPath []string

func Final() {
	for i := range roomPaths {
		if roomPaths[i] != "" {
			finalPath = append(finalPath, roomPaths[i])
		}
	}
}

// place rooms in grid
// func grid() {
// 	var grid [30][30]string
// 	for row := 0; row < len(grid); row++ {
// 		for column := 0; column < len(grid); column++ {
// 			grid[row][column] = " "
// 		}
// 	}
// 	for i := range roomList {
// 		for row := 0; row < len(grid); row++ {
// 			for column := 0; column < len(grid); column++ {
// 				if row == roomList[i].row-1 && column == roomList[i].column-1 {
// 					grid[row][column] = "[" + roomList[i].name + "]"
// 				}
// 			}
// 		}
// 	}
// 	for i := range grid {
// 		for _, ele := range grid[i] {
// 			fmt.Print(ele)
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	getAnts()
	getRooms()
	linkRooms()
	assignStart()
	assignEnd()
	allPaths(Start)
	Final()
	fmt.Println(finalPath)
}
