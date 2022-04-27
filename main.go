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
	noOfAnts int
}

type Ant struct {
	name     string
	room     *room
	prevRoom *room
	path     string
}

var ants []*Ant

//This is to obtain and create an address for n number of ants.
func getAnts() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	getants := bufio.NewScanner(data)
	line := 0

	for getants.Scan() {
		if getants.Text() == "" {
			fmt.Println("ERROR: invalid data format")
			log.Fatal()
		}
		line++
		if line == 1 {
			a, err2 := strconv.Atoi(getants.Text())
			if a == 0 {
				fmt.Println("ERROR: invalid data format")
				fmt.Println("No ants found")
				log.Fatal()
			}
			if err2 != nil {
				fmt.Println("ERROR: invalid data format")
				fmt.Println("No ants found")
				log.Fatal()
			}
			ants = make([]*Ant, a)
			for i := 0; i < a; i++ {
				antName := &Ant{name: strconv.Itoa(i + 1)}
				ants[i] = antName
			}
		}
	}
}

var roomList []*room

// to initialise rooms with their own address
func getRooms() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}

	var emptyString string
	var getCoOrd string
	line := 0
	getCoOrds := bufio.NewScanner(data)
	for getCoOrds.Scan() {
		line++
		if line > 1 {
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
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
	var emptyString string
	var links []string
	line := 0
	linksInfo := bufio.NewScanner(data)
	for linksInfo.Scan() {
		line++
		if line > 1 {
			if strings.Contains(linksInfo.Text(), "-") {
				emptyString = linksInfo.Text()
				links = append(links, emptyString)
			} else {
				emptyString = ""
			}
		}
	}

	for i := range links {
		for j := range links[i] {
			if links[i][j] == '-' {
				linkString := strings.Split(links[i], "-")
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
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
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
	if Start == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No Start Room Found")
		log.Fatal()
	}
	lenStart = len(Start.nextRoom)
	if lenStart == 0 {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No rooms connecting to Start")
		log.Fatal()
	}
	roomPaths = make([]string, 5)
}

// function to assign the end room for ants
var End *room

func assignEnd() {
	data, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("File Error")
		log.Fatal()
	}
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
	if End == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No End Room assigned")
		log.Fatal()
	}
	if End.nextRoom == nil {
		fmt.Println("ERROR: invalid data format")
		fmt.Println("No rooms connecting to end")
		log.Fatal()
	}
	End.noOfAnts = len(ants)
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

	for k := range roomList {
		if roomList[k].nextRoom != nil {
			roomList[k].visited = 0
			roomList[k].nextRoom = nil
		}
	}

	for i := range roomPaths {
		b := strings.Split(roomPaths[i], ",")
		for o := 0; o < len(b)-1; o++ {
			for k := range roomList {
				for l := range roomList {
					if b[o] == roomList[k].name && b[o+1] == roomList[l].name {
						roomList[k].nextRoom = append(roomList[k].nextRoom, roomList[l])
					}
				}
			}
		}
	}
}

// this sorts the elements of roomPaths in ascending order
func Sort() {
	for i := 0; i < len(finalPath)-1; i++ {
		if len(finalPath[i]) > len(finalPath[i+1]) {
			g := finalPath[i]
			h := finalPath[i+1]
			finalPath[i] = h
			finalPath[i+1] = g
		} else if len(finalPath[i+1]) < len(finalPath[i]) {
			g := finalPath[i+1]
			h := finalPath[i]
			finalPath[i] = h
			finalPath[i+1] = g
		}
	}
}

// this returns the index of the smallest path
var roomLength []int

func minPath(roomLength []int) int {
	min := roomLength[0]
	index := 0
	for i, room := range roomLength {
		if room < min {
			min = room
			index = i
		}
	}
	return index
}

// this assigns the appropriate path to each ant with the smallest number of turns
func assignPaths() {
	for i := range finalPath {
		a := strings.Split(finalPath[i], ",")
		roomLength = append(roomLength, (len(a) - 2))
	}
	for n := range ants {
		ants[n].path = finalPath[minPath(roomLength)]
		roomLength[minPath(roomLength)]++
		a := strings.Split(ants[n].path, ",")
		for _, ele := range Start.nextRoom {
			if ele.name == a[1] {
				ants[n].room = ele
			}
		}
	}
}

var antPath string

func TraversePath(r *room) {
	endAnts := 0
	for endAnts != End.noOfAnts {
		for i := range ants {
			if i < len(ants)-1 {
				if ants[i].room.visited == 0 && !ants[i].room.end {
					ants[i].room.visited = 1
					antPath += string("L"+ants[i].name+"-"+ants[i].room.name) + " "
					ants[i].prevRoom = ants[i].room
					ants[i].room = ants[i].room.nextRoom[0]
				} else if ants[i].room.visited == 0 && ants[i].room.end {
					if !strings.Contains(antPath, string("L"+ants[i].name)+"-"+ants[i].room.name) {
						antPath += "L" + ants[i].name + "-" + ants[i].room.name + " "
						ants[i].prevRoom = ants[i].room
						endAnts++
					}
				}
			} else if i == len(ants)-1 {
				if ants[i].room.end {
					antPath += string("L"+ants[i].name+"-"+ants[i].room.name) + " "
					endAnts++
				}
				if ants[i].room.visited == 0 && !ants[i].room.end {
					ants[i].room.visited = 1
					antPath += "L" + ants[i].name + "-" + ants[i].room.name + " "
					ants[i].prevRoom = ants[i].room
					ants[i].room = ants[i].room.nextRoom[0]
				}
				for j := range ants {
					if ants[j].prevRoom != nil {
						ants[j].prevRoom.visited = 0
					}
				}
			}
		}

		antPath += "\n"
	}

	antPath = antPath[:len(antPath)-1]
}

func main() {
	getAnts()
	getRooms()
	linkRooms()
	assignStart()
	assignEnd()
	allPaths(Start)
	Final()
	Sort()
	assignPaths()
	TraversePath(Start)
	file, _ := os.ReadFile(os.Args[1])
	fmt.Println(string(file) + "\n")
	fmt.Println(antPath)
}
