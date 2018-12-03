package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	cloth := make([][]string,1000)
	for y := range cloth {
		cloth[y] = make([]string,1000)
	}

	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	var conflict int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		specs := strings.Split(l, " ")
		id := string(specs[0][1:])
		pos := strings.Split(specs[2], ",")
		xpos,_ := strconv.Atoi(pos[0])
		ypos,_ := strconv.Atoi(strings.Split(pos[1],":")[0])
		xsize,_ := strconv.Atoi(strings.Split(specs[3], "x")[0])
		ysize,_ := strconv.Atoi(strings.Split(specs[3], "x")[1])


		for xidx := xpos; xidx < xpos+xsize; xidx++ {
			for yidx := ypos; yidx < ypos+ysize; yidx++ {
				if cloth[xidx][yidx] == "" {
					cloth[xidx][yidx] = id
				} else {
					cloth[xidx][yidx] = "O"
				}
			}
		}


	}


	for _,x := range cloth {
		for _,y := range x {
			if y == "O" {
				conflict = conflict + 1
			}
		}
	}

fmt.Println(conflict)
}
