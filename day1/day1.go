package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}

	scanner := bufio.NewScanner(f)
	var freq int
	for scanner.Scan() {
		l := scanner.Text()
		op := string(l[0])
		val,_ := strconv.Atoi(l[1:])
		if op == "+" {
			freq = freq + val
		} else if op == "-" {
			freq = freq - val
		}
	}

	fmt.Println(freq)

}
