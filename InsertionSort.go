import (
	"fmt"
)

func InsertionSort(A []int) {
	var key int
	var i int
	for j := 1; j < len(A); j++ {
		key = A[j]
		i = j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	// return A
}

func main() {
	a := []int{10, 5, 7, 9, 7, 5, 6, 2, 3, 1}
	InsertionSort(a)
	fmt.Println(a)
}
