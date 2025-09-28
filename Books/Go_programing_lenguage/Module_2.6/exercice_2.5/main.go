package main

import "fmt"


// PopCountClearRightmost cuenta los bits usando x & (x - 1).
func PopCountClearRightmost(x uint64) int {
    count := 0
    for x != 0 {
        x = x & (x - 1)
        count++
    }
    return count
}

func main() {
	fmt.Println(PopCountClearRightmost(123456))
}
