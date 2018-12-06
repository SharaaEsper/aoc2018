package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)


func up(x, y int) (int, int) {
	return x, y-1
}
func down(x, y int) (int, int) {
	return x, y+1
}
func left(x, y int) (int, int) {
	return x-1, y
}
func right(x, y int) (int, int) {
	return x+1, y
}
func get_coords(coord string)(int, int) {
	a := strings.Split(coord, ",")
	x,_ := strconv.Atoi(a[0])
	y,_ := strconv.Atoi(strings.TrimSpace(a[1]))
	return x, y
}
func get_distance(x1, y1, x2, y2 int) (int) {
	df := math.Abs(float64(x1 - x2)) + math.Abs(float64(y1 - y2))
	di := int(df)
	return di
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}


	var_key := make(map[int]string,0)
	idx := 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		var_key[idx] =  l
		idx = idx + 1
	}


	grid := make([][]int,1000)
	for i :=0; i < 1000;i++ {
			grid[i] = make([]int,1000)
	}




	for k,v := range var_key {
		a := strings.Split(v, ",")
		a[1] = strings.TrimSpace(a[1])
	  x,_ := strconv.Atoi(a[0])
		y,_ := strconv.Atoi(a[1])
		grid[x][y] = k
	}



	//Create bounds
	var leftbound [2]int
	leftbound[0] = 9999
	leftbound[1] = 0
	var rightbound [2]int
	rightbound[0] = 0
	rightbound[1] = 0
	var topbound [2]int
	topbound[0] = 0
	topbound[1] = 9999
	var bottombound [2]int
	bottombound[0] = 0
	bottombound[1] = 0
	for _,v := range var_key {
		x,y := get_coords(v)
		if x < leftbound[0] {
			leftbound[0] = x
			leftbound[1] = y
		}
		if x > rightbound[0] {
			rightbound[0] = x
			rightbound[1] = y
		}
		if y < topbound[1] {
			topbound[0] = x
			topbound[1] = y
		}
		if y > bottombound[1] {
			bottombound[0] = x
			bottombound[1] = y
		}
	}


	//Lets Fill in The box with [70] to be able to just check for [0] if they're OOB
	for x := leftbound[0]; x < rightbound[0]; x++ {
		for y := topbound[1]; y < bottombound[1]; y++ {
			grid[x][y] = 70
		}
	}

	//Now lets overwrite that with some comparisons between nodes
	leftnode := 1
	rightnode := 1
	upnode := 1
	downnode := 1
	for node := 1; node < len(var_key) +1; node ++ {
		x,y := get_coords(var_key[node])
		//Dumb search, we find "closest" in each direction
		leftnode = 1
		rightnode = 1
		upnode = 1
		downnode = 1
		var cx,cy int
		for k,v := range var_key {
			cx,cy = get_coords(v)
			if cx < x {
				leftnode = k
			}
			if cx > x {
				rightnode = k
			}
			if cy < y {
				upnode = k
			}
			if cy > y {
				downnode = k
			}
		}
		//Now we figure out the distance between them
		leftx,lefty := get_coords(var_key[leftnode])
		leftdistance := get_distance(x, y, leftx, lefty)
		rightx,righty := get_coords(var_key[rightnode])
		rightdistance := get_distance(x, y, rightx, righty)
		upx,upy := get_coords(var_key[upnode])
		updistance := get_distance(x, y, upx, upy)
		downx,downy := get_coords(var_key[downnode])
		downdistance := get_distance(x, y, downx, downy)

		//
		fmt.Println(leftdistance,rightdistance,updistance,downdistance)
		fmt.Println(leftnode,rightnode)

	}
}
