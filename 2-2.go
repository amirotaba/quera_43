package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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


func Read(x int, times map[int]int) map[int]int {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < x; i++{
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		stext := strings.Split(text, " ")
		times[Parse(stext[1])] = Parse(stext[2])
	}
	return times
}

func common(times map[int]int) []int {
	var common int
	starts := make([]int, 0)
	commons := make([]int, 0)
	for i := range times {
		starts = append(starts, i)
	}
	sort.Ints(starts)
	for u := 0; u < len(starts)-1; u++{
		if times[starts[u]] > starts[u+1] {
			if times[starts[u+1]] < times[starts[u]] {
				common = times[starts[u+1]] - starts[u+1]
				commons = append(commons, common)
			}else {
				common = times[starts[u]] - starts[u+1]
				commons = append(commons, common)
			}
		}
	}
	return commons
}

func Sum(list []int) int {
	sum := 0
	for u := range list {
		sum += list[u]
	}
	return sum
}

func main() {
	var n, m int
	times := make(map[int]int, n+m)
	fmt.Scan(&n)
	times = Read(n, times)
	fmt.Scan(&m)
	times = Read(m, times)
	commons := common(times)
	sum := Sum(commons)
	fmt.Println(sum)
}