package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

func main() {


	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	data := make([]string,0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	sort.Strings(data)

	guard_key := make(map[string][]string,0)

	var t1 string
	var t2 string
	var current_gid string
	tr,_ := regexp.Compile("\\[.*\\]")
	cg,_ := regexp.Compile("Guard #.+ begins")
	woker,_ := regexp.Compile("wakes up")
	for _,l := range data {
		t2 = t1
		t1 = tr.FindString(l)
		t1 = t1[1:len(t1)-1]
		gid := cg.FindString(l)
		gid2 := strings.Split(gid, " ")
		if  len(gid2) > 1  {
			if len(guard_key[gid2[1][1:]]) <1 {
				guard_key[gid2[1][1:]] = append(guard_key[gid2[1][1:]] ,"")
			}
			current_gid = gid2[1][1:]
			continue
		}
		if woker.MatchString(l) {
			guard_key[current_gid] = append(guard_key[current_gid],t2)
			guard_key[current_gid] = append(guard_key[current_gid],t1)
		}
	}

	minutes := make(map[string][]int,60)
	for gid,_ := range guard_key {
		for init := 0; init < 60; init++ {
			minutes[gid] = append(minutes[gid],0)
		}
	}
	for i := 0; i < 60; i++ {
		for gid,_ := range guard_key {
			for k,_ := range guard_key[gid] {
				if k > 0 && k % 2 == 0 {
					st1,_ := strconv.Atoi(guard_key[gid][k-1][len(guard_key[gid][k-1])-2:])
					st2,_ := strconv.Atoi(guard_key[gid][k][len(guard_key[gid][k])-2:])
					if i >= st1 && i < st2 {
						minutes[gid][i] = minutes[gid][i] + 1
					}
				}
			}
		}
	}


	var mostslept string
	var minuteslept int
	var sleepcounter int
	for i,_ := range minutes {
		for k,m := range minutes[i] {
			if m > sleepcounter {
				minuteslept = k
				mostslept = i
				sleepcounter = m
			}
		}
	}

	sleepyguardi,_ := strconv.Atoi(mostslept)
	fmt.Println("Asshole sleepy guard ID is:", sleepyguardi)
	fmt.Println("He slept the most of everyone on minute:", minuteslept)
	fmt.Println("The solution is:", sleepyguardi * minuteslept)
}
