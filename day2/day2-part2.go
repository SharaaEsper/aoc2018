package main

import (
	"fmt"
	"bufio"
	"os"
)


func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(f)
	ids := make([]string,0)
	for scanner.Scan() {
		l := scanner.Text()
		ids = append(ids, l)
	}

	var solution string
	var line1 string
	var line2 string
	for index,line := range ids {
		line1 = line
		for _,tline := range ids[index+1:] {
			var misscount int
			for i :=0; i <= len(line) - 1; i++ {
				if line[i] != tline[i] {
					misscount = misscount + 1
				}
				if misscount > 1 {
					break
				}
			}
			if misscount == 1 {
				line2 = tline
				break
			}
		}
		if line2 != "" {
			break
		}
	}
	for ci,chr := range line1 {
		if string(chr) == string(line2[ci]) {
			solution = join(solution, string(chr))
		}
	}
	fmt.Println(solution)

}
