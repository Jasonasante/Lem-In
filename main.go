package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"container/list"
)

type room struct {
	name     string
	column   int
	row      int
	nextRoom map[string]*room
	start    bool
	end      bool
	visited  int
}

var (
	roomList []*room
	Start     *room
)
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
		if roomList[i].start == true {
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

	for _,ele:=range roomList{
		ele.nextRoom=make(map[string]*room,0)
	}
	
	for i := range links {
		for j := range links[i] {
			if links[i][j] == '-' {
				for k := range roomList {
					for o := range roomList {
						if string(links[i][j-1]) == roomList[k].name && roomList[o].name == string(links[i][j+1]) {
							roomList[k].nextRoom[roomList[o].name] = roomList[o]
						} else if string(links[i][j+1]) == roomList[k].name && roomList[o].name == string(links[i][j-1]) {
							roomList[k].nextRoom[roomList[o].name] = roomList[o]
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
// var (
// 	roomPaths [][]*room
// 	count     int
// )


func pathRec(r *room) {
	// rooms:=r
	nextRoom := r.nextRoom
	// fmt.Println(nextRoom)
	// fmt.Println("start= ",r.name)
	fmt.Println()
	for _, ele := range nextRoom {
		fmt.Print("  check= ", ele.name)
		fmt.Print(" visited:= ", ele.visited)

	}
	fmt.Println()

	for i := range nextRoom {
		if !nextRoom[i].end && nextRoom[i].visited == 0 {
			for k := range nextRoom[i].nextRoom {
				if !nextRoom[i].nextRoom[k].end && nextRoom[i].nextRoom[k].visited == 0 {
					nextRoom[i].visited = 1
					pathRec(nextRoom[i])
				}
			}
			nextRoom[i].visited = 1
		} else if nextRoom[i].end {
			fmt.Println("end")
			pathRec(Start)

		}
	}
}

func RouteToEnd() {
	// roomPaths := make([][]*room, 20)
	// count := 0
	for i := range roomList {
		if roomList[i].start == true {
			Start = roomList[i]
		}
	}
	Start.visited = 1
	pathRec(Start)

	// nextRoom := Start.nextRoom

	// // for j:=range nextRoom[1].nextRoom{
	// // 	fmt.Println(nextRoom[1].nextRoom[j])
	// // }
	// roomPaths[count] = append(roomPaths[count], Start)
	// for i := range nextRoom {
	// 	Start.visited = 1
	// 	if nextRoom[i].end {
	// 		roomPaths[count] = append(roomPaths[count], nextRoom[i])
	// 		fmt.Println("end", nextRoom[i])
	// 		count++
	// 		nextRoom = Start.nextRoom
	// 	} else if !nextRoom[i].end && nextRoom[i].visited == 0 {
	// 		roomPaths[count] = append(roomPaths[count], nextRoom[i])
	// 		// fmt.Println("check", nextRoom[i].name)
	// 		nextRoom[i].visited = 1
	// 		nextRoom = nextRoom[i].nextRoom
	// 	}
	// }

	// for _, ele := range roomPaths[0] {
	// 	fmt.Println(ele)
	// }
	// fmt.Println(roomPaths)
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
	RouteToEnd()
}