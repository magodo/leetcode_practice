package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	m := map[int]struct{}{}
	for _, i := range A {
		m[i] = struct{}{}
	}
	for i := 1; i <= 100000; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 100001
}
