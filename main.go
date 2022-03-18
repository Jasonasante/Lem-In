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
	// fmt.Println("links:= ", links)
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
	// fmt.Println(lenStart)
	// for _,ele:=range Start.nextRoom{
	// 	fmt.Println(*ele)
	// }
	roomPaths = make([]string, lenStart+1)
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
			ele.nextRoom = nil
			End = ele
		}
	}
}

// find path from start to end
var (
	count      int
	roomPaths  []string
	potPath    string
	returnPath string
	visitedNR  int
)

func allPaths(r *room) {
	prevRoom := r
	nextRoom := r.nextRoom
	fmt.Println(roomPaths)
	counter := 0

	// fmt.Println()
	// fmt.Println("rNameSLice last index:= ",rNamesSlice[len(rNamesSlice)-2])

	if prevRoom.end {
		roomPaths[count] += prevRoom.name
		verifyPath(roomPaths[count])
		Start.visited = 0
		End.visited = 0
		count++
		allPaths(Start)
	}
	for _, ele := range nextRoom {
		if ele.visited == 1 {
			counter++
		}
	}
	if counter == len(nextRoom) {
		prevRoom.visited = 1
		dEndNameSlice := strings.Split(roomPaths[count], ",")
		for _, room := range roomList {
			if len(dEndNameSlice) >= 2 {
				if dEndNameSlice[len(dEndNameSlice)-2] == room.name {
					for j := range room.nextRoom {
						if room.nextRoom[j].visited == 0 && room.nextRoom[j].name != prevRoom.name {
							allPaths(room.nextRoom[j])
						}
					}
				}
				
			}

		}
		for i := range nextRoom {
			if nextRoom[i] == Start {
				fmt.Println("stop")
				// roomPaths = append(roomPaths[:count],roomPaths[:count-1]...)
				roomPaths[count] = ""
				fmt.Println("why")
				allPaths(Start)
			}
		}
	} else {
		for i, rooms := range nextRoom {
			if count < lenStart {
				rNamesSlice := strings.Split(roomPaths[count], ",")
				if !contains(rNamesSlice, rooms.name) && (rooms.visited == 0) {
					roomPaths[count] += prevRoom.name + ","
					prevRoom.visited = 1
					allPaths(nextRoom[i])
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

// checks for other paths
func otherPaths(r *room) {
	prevRoom := r
	nextRoom := r.nextRoom
	if prevRoom.end {
		returnPath += prevRoom.name
	}
	for _, ele := range nextRoom {
		if ele.visited == 1 {
			visitedNR++
		}
	}

	if visitedNR == len(nextRoom) {
		returnPath = potPath
		visitedNR = 0
	} else {
		for i, rooms := range nextRoom {
			returnPathSlice := strings.Split(returnPath, ",")
			if !contains(returnPathSlice, rooms.name) && (rooms.visited == 0) {
				returnPath += prevRoom.name + ","
				prevRoom.visited = 1
				otherPaths(nextRoom[i])
			}
		}
	}
}

// verifies if the path we collect is the shortest option
func verifyPath(s string) {
	potPath = s
	potPlathSlice := strings.Split(potPath, ",")
	fmt.Println(potPlathSlice)
	var sent *room
	// so with example1 returnPath may be "1,3," or "1,2,"
	var next *room
	for i := 0; i < 2; i++ {
		returnPath += potPlathSlice[i] + ","
	}
	// fmt.Println("potPlathSlice check", potPlathSlice)
	// fmt.Println("returnPath--", returnPath)
	// If the incoming string/path starts with the the start room and ends with an end room,
	// then check if it is the shortest part available. Else make that roomPath[Count]=="".
	if (potPlathSlice[0] == Start.name) && (potPlathSlice[(len(potPlathSlice)-1)] == End.name) {
		// the room at s[4] is the nextroom of start's nextroom. So we make that room visited
		// and make the sent variable s[2].
		if len(potPlathSlice) != 2 {
			for _, roomele := range roomList {
				for i := 2; i < len(potPlathSlice)-1; i++ {
					if potPlathSlice[i] == roomele.name {
						roomele.visited = 0
					}
				}
				if potPlathSlice[2] == roomele.name {
					roomele.visited = 1
				}
				if potPlathSlice[1] == roomele.name {
					sent = roomele
				}
			}
			// Going through sent's nextroom, check if that room is not visited (which in this case s[4] would have been).
			// If that room is not visited, make sent ==ele and carry out the otherPaths(sent) (which is similar to the allPaths()
			// but it appends the path generated to returnPath).
			// If the returnPath is longer than or equal to potPath, then make revert returnPath.
			// If the returnPath is shorter that potPath, make potPath= returnPath and repeat until all the next rooms have been visited.
			// Once all rooms have been visited, make all rooms not visited, then all the rooms in potPath visited, then finally; roomPath[count]==potPath.
			for _, ele := range sent.nextRoom {
				if ele.visited == 0 {
					next = ele
					otherPaths(next)
					if len(returnPath) > len(potPath) {
						returnPathSlice := strings.Split(returnPath, ",")
						for _, roomele := range roomList {
							for i := 3; i < len(returnPathSlice)-1; i++ {
								if returnPathSlice[i] == roomele.name {
									roomele.visited = 0
								}
							}
							if returnPathSlice[3] == roomele.name {
								roomele.visited = 1
							}
						}
						returnPath = ""
						for i := 0; i < 2; i++ {
							returnPath += returnPathSlice[i] + ","
						}
					} else {
						returnPathSlice := strings.Split(returnPath, ",")
						fmt.Println("less than returnPath", returnPathSlice)
						for _, roomele := range roomList {
							roomele.visited = 0
						}
						for i := range returnPathSlice {
							for _, roomele := range roomList {
								if returnPathSlice[i] == roomele.name {
									roomele.visited = 1
								}
							}
						}

						roomPaths[count] = returnPath
						returnPath = ""
					}
				}
			}
		}
	} else {
		roomPaths[count] = ""
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
	getRooms()
	linkRooms()
	assignStart()
	assignEnd()
	allPaths(Start)
	// for i := range roomList {
	// 	for _, nextroom := range roomList[i].nextRoom {
	// 		fmt.Print(roomList[i].name)
	// 		fmt.Print(":= ")
	// 		fmt.Println(nextroom)
	// 	}
	// }
	fmt.Println("RoomPaths:=")
	fmt.Println(roomPaths)
}
