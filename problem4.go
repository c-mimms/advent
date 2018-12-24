package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func solve1() {
	timestamps := readFileToArray("input4.txt")
	sort.Strings(timestamps)

	currentGuard := 0
	asleepMin := 0
	asleepTimes := make(map[int][60]int)

	for _, time := range timestamps {
		//Set the current guard
		if strings.Contains(time, "Guard") {
			re := regexp.MustCompile("#\\d+")
			guardNum := re.FindString(time)
			currentGuard, _ = strconv.Atoi(strings.Trim(guardNum, "#"))
		} else if strings.Contains(time, "falls") {
			asleepMin, _ = strconv.Atoi(strings.Split(time, ":")[1][:2])
		} else if strings.Contains(time, "wakes") {
			//Get just the minutes part
			wakeMin, _ := strconv.Atoi(strings.Split(time, ":")[1][:2])

			guardTimes := asleepTimes[currentGuard]
			for i := asleepMin; i < wakeMin; i++ {
				guardTimes[i] += 1
			}
			asleepTimes[currentGuard] = guardTimes
		}
	}

	mostSleep := 0
	mostGuard := 0
	mostMinute := -1
	for guard, times := range asleepTimes {
		sleep := 0
		guardMostSleepMin := -1
		mostMin := 0
		for min, num := range times {
			sleep += num
			if num > guardMostSleepMin {
				mostMin = min
				guardMostSleepMin = num
			}
		}
		if sleep > mostSleep {
			mostSleep = sleep
			mostGuard = guard
			mostMinute = mostMin
		}
	}
	printVal(mostMinute * mostGuard)
}

func solve2() {
	timestamps := readFileToArray("input4.txt")
	sort.Strings(timestamps)

	currentGuard := 0
	asleepMin := 0
	asleepTimes := make(map[int][60]int)

	for _, time := range timestamps {
		//Set the current guard
		if strings.Contains(time, "Guard") {
			re := regexp.MustCompile("#\\d+")
			guardNum := re.FindString(time)
			currentGuard, _ = strconv.Atoi(strings.Trim(guardNum, "#"))
		} else if strings.Contains(time, "falls") {
			asleepMin, _ = strconv.Atoi(strings.Split(time, ":")[1][:2])
		} else if strings.Contains(time, "wakes") {
			//Get just the minutes part
			wakeMin, _ := strconv.Atoi(strings.Split(time, ":")[1][:2])

			guardTimes := asleepTimes[currentGuard]
			for i := asleepMin; i < wakeMin; i++ {
				guardTimes[i] += 1
			}
			asleepTimes[currentGuard] = guardTimes
		}
	}

	mostTimesSleep := 0
	mostTimesSleepGuard := 0
	mostTimesSleepMin := 0
	for guard, asleepTime := range asleepTimes {
		for min, times := range asleepTime {
			if times > mostTimesSleep {
				mostTimesSleep = times
				mostTimesSleepGuard = guard
				mostTimesSleepMin = min
			}
		}
	}
	printVal(mostTimesSleepGuard * mostTimesSleepMin)
	printVal(mostTimesSleep)
}
