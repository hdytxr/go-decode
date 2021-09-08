package netflix

import (
	"fmt"
	"strconv"
)

func GroupTitles(strs []string) [][]string {
	var output [][]string
	if len(strs) == 0 {
		return output
	}

	res := make(map[string][]string)

	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			index := c - 'a'
			count[index]++
		}

		key := ""
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		res[key] = append(res[key], s)
	}

	for _, v := range res{
		output = append(output, v)
	}

	return output
}

func GroupTest() {
	titles := []string{"duel","dule","speed","spede","deul","cars"}
	query := "spede"
	output := GroupTitles(titles)

	for i, o := range output{
		// if find(o, query) {
		// 	fmt.Println(o)
		// }
		if o[i] == query {
			fmt.Println(o)
		}
	}
}