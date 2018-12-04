package main

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"
	"time"
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

	var t1parsed time.Time
	var t2parsed time.Time
	var zero_dur time.Duration
	layout := "2006-01-02 15:04"
	guard_sleep := make(map[string]time.Duration)
	for gid,guard := range guard_key {
		guard_sleep[gid] = zero_dur
		for k,timestamp := range guard {
			if int(k) % 2 == 0 && int(k) > 0{
				t2parsed,_ = time.Parse(layout, timestamp)
				guard_sleep[gid] = guard_sleep[gid] + t2parsed.Sub(t1parsed)
			} else {
				t1parsed,_ = time.Parse(layout, timestamp)
			}
		}
	}

	var sleepyguard string
	for gid,_ := range guard_sleep {
		if sleepyguard == "" {
			sleepyguard = gid
		} else {
			if guard_sleep[gid] > guard_sleep[sleepyguard] {
				sleepyguard = gid
			}
		}
	}


	minutes := make([]int,60)
	for i := 0; i < 60; i++ {
		for k, _ := range guard_key[sleepyguard] {
			if k > 0 && k % 2 == 0 {
				st1,_ := strconv.Atoi(guard_key[sleepyguard][k-1][len(guard_key[sleepyguard][k-1])-2:])
				st2,_ := strconv.Atoi(guard_key[sleepyguard][k][len(guard_key[sleepyguard][k])-2:])
				if i >= st1 && i < st2 {
					minutes[i] = minutes[i] + 1
				}
			}
		}
	}

	var mostslept int
	for k,m := range minutes {
		if k == 0 {
			mostslept = k
		} else {
			if m > minutes[mostslept] {
				mostslept = k
			}
		}
	}

	sleepyguardi,_ := strconv.Atoi(sleepyguard)
	fmt.Println("Asshole sleepy guard ID is:", sleepyguardi)
	fmt.Println("He slept the most on minute:", mostslept)
	fmt.Println("The solution is:", sleepyguardi * mostslept)
}
