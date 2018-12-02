package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	var doubles int
	var triples int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		var has_double bool
		var has_triple bool
		for i :=0; i <= len(l) - 1; i++ {
			chrcount := strings.Count(l, string(l[i])) 
			if chrcount == 2 {
				has_double = true
			} else if chrcount == 3 {
				has_triple = true
			}
		}
		if has_double {
			doubles = doubles + 1
		}
		if has_triple {
			triples = triples + 1
		}
	}
	fmt.Println(doubles * triples)
}
