package main

import "fmt"


// PopCountShift cuenta los bits usando desplazamiento y prueba de bits.
func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func main() {
	fmt.Println(PopCountShift(123456))
}
