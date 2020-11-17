package main

import (
	"fmt"
)

func Merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	var L []int
	var R []int
	var k int = p
	var i int
	var j int
	for i = 0; i < n1; i++ {
		L = append(L, A[k])
		k += 1
	}
	for j = 0; j < n2; j++ {
		R = append(R, A[k])

		k += 1
	}
	L = append(L, 99999999)
	R = append(R, 99999999)
	k = p
	i = 0
	j = 0
	for k = p; k <= r; k++ {
		if L[i] >= R[j] {
			A[k] = R[j]
			j += 1

		} else {
			A[k] = L[i]
			i += 1
		}

	}
	fmt.Println(L, R)
}
func MergeSort(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)

	}

}

func main() {
	a := []int{10, 5, 7, 9, 7, 5, 6, 2, 3, 1}
	MergeSort(a, 0, 9)
	fmt.Println(a)

}

// 2ND WAY
// func merge(a, b []int) []int {
// 	size, i, j := len(a)+len(b), 0, 0
// 	result := make([]int, size)
// 	for k := 0; k < size; k++ {
// 		switch true {
// 		case i == len(a):
// 			//all the elements of a already been judged
// 			//assuming a and b both are sorted, this happens because
// 			//some cases will have not equal a and b, so it might
// 			// be a possibility that one array got finished earlier.
// 			result[k] = b[j]
// 			j += 1
// 		case j == len(b):
// 			//alll the elements of a already been judged
// 			//assuming a nd b both are sorted
// 			result[k] = a[i]
// 			i += 1
// 		case a[i] > b[j]:
// 			result[k] = b[j]
// 			//increment the index of b array because its present index
// 			//is already appended to the result array
// 			j += 1
// 		case a[i] < b[j]:
// 			//increment the index of a array because its present index
// 			//element is already appended to the result array
// 			result[k] = a[i]
// 			i += 1
// 		case a[i] == b[j]:
// 			//in case of equality
// 			result[k] = b[j]
// 			j += 1
// 		}
// 	}
// 	return result
// }
// func MergeSort(numbers []int) []int {
// 	if len(numbers) < 2 {
// 		return numbers
// 	}
// 	middle := int(len(numbers) / 2)
// 	return merge(MergeSort(numbers[middle:]), MergeSort(numbers[:middle]))
// }
// func main() {
// 	a := []int{2, 1, 3, 4, 50, 78, 32, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
// 	fmt.Println(MergeSort(a))
// }
