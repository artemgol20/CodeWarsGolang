// https://www.codewars.com/kata/5a662a02e626c54e87000123/solutions/go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(ExtraPerfect(n))

}

func ExtraPerfect(n int) []int {
	var array []int
	for i := 1; i <= n; i++ {
		var binary = strconv.FormatInt(int64(i), 2)
		if getChar(binary, 0) == getChar(binary, len([]rune(binary))-1) {
			array = append(array, i)
		}
	}
	return array
}

func getChar(str string, index int) rune {
	return []rune(str)[index]
}
