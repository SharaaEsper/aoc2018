package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	f, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}

/*
97 = a
122 = z
65 = A
90 = Z
Uppercase -> Lowercase always is always +32
*/



	data := make([]byte, len(f))
	del_idx := 0
	current_idx := 0
	count := make([]int,0)

	for d := 65; d < 91; d++ {
		data = make([]byte, len(f))
		copy(data,f)
		del_idx = 0
		for {
			if del_idx > len(data)-2 {
				break
			}
			if data[del_idx] == byte(d) || data[del_idx] == byte(d+32) {
				data = append(data[:del_idx], data[del_idx+1:]...)
				continue
			}
			del_idx = del_idx + 1
		}
		current_idx = 0
		var edited bool
		for {
			if current_idx > len(data)-2 {
				if edited {
					current_idx = 0
					edited = false
				} else {
					count = append(count,len(data))
					break
				}
			}
			if data[current_idx] == data[current_idx+1] - byte(32) || data[current_idx] == data[current_idx+1] + byte(32) {
				data = append(data[:current_idx], data[current_idx+2:]...)
				edited = true
				continue
			}
			if current_idx > 0 {
				if data[current_idx] == data[current_idx-1] - byte(32) || data[current_idx] == data[current_idx+1] + byte(32) {
					data = append(data[:current_idx-1], data[current_idx+1:]...)
					edited = true
					continue
				}
			}
			current_idx = current_idx +1
		}
	}

	smolboi := 50000
	for _,v := range count {
		if v < smolboi {
			smolboi = v
		}
	}
	fmt.Println("Smalboi is:", smolboi)


}
