package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func Parse(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	} else {
		return -10000
	}
}

func Meeting(n int, times [][]int) [][]int {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++{
		time := make([]int, 0)
		text ,_ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		stext := strings.Split(text, " ")
		for o := 0; o < len(stext); o++{
			time = append(time, Parse(stext[o]))
		}
		times = append(times, time)
	}
	return times
}

func FindDur() int{
	reader := bufio.NewReader(os.Stdin)
	text ,_ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	stext := strings.Split(text, " ")
	duration := Parse(stext[1])
	return duration
}

func Free(times [][]int, duration int) []int {
	output := make([]int, 0)
	if times[0][1] - duration >= 0 {
		output = append(output, 0, duration)
		return output
	}else{
		for i := range times {
			if i < len(times) {
				if times[i+1][1] - times[i][2] >= duration {
					start := times[i][2]
					output = append(output, start, start+duration)
					return output
				}
			}else {
				start := times[len(times)][2]
				output = append(output, start, start+duration)
				return output
			}
		}
	}
	return nil
}

func main() {
	var n int
	times := make([][]int, 0)
	fmt.Scan(&n)
	free := Free(Meeting(n, times), FindDur())
	fmt.Println(free)
}
