package main

import (
	"fmt"
	"io/ioutil"
)

func main() {



/*
97 = a
122 = z
65 = A
90 = Z
Uppercase -> Lowercase always is always +32
*/


	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}


	current_idx := 0
	var edited bool
	for {
		if current_idx > len(data)-2 {
			if edited {
				current_idx = 0
				edited = false
			} else {
				break
			}
		}
		if data[current_idx] == data[current_idx+1] - 32 || data[current_idx] == data[current_idx+1] + 32 {
			data = append(data[:current_idx], data[current_idx+2:]...)
			edited = true
			continue
		}
		if current_idx > 0 {
			if data[current_idx] == data[current_idx-1] - 32 || data[current_idx] == data[current_idx+1] + 32 {
				data = append(data[:current_idx-1], data[current_idx+1:]...)
				edited = true
				continue
			}
		}
		current_idx = current_idx +1
	}



	fmt.Println(len(data))

}
