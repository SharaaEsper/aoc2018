package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	var matched bool
	var freq int
	prevvalues := make(map[int]int)
	for {
		f, err := os.Open("input")
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l := scanner.Text()
			op := string(l[0])
			val,_ := strconv.Atoi(l[1:])
			if op == "+" {
				freq = freq + val
			} else if op == "-" {
				freq = freq - val
			}
			if _,exists := prevvalues[freq]; exists {
				matched = true
				break
			} else {
				prevvalues[freq] = freq
			}
		}
		if matched {
			break
		}
	}
	fmt.Println(freq)
}
