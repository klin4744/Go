package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		numbers []int
		target  int
		answer  bool
	}
	tests := []test{
		test{
			numbers: []int{1, 2, 3, 4, 5, 6},
			target:  7,
			answer:  true,
		},
		test{
			numbers: []int{1, 9, 3, 3, 5, 6},
			target:  6,
			answer:  true,
		},
		test{
			numbers: []int{1, 2, 9, 23, 0, 10},
			target:  100,
			answer:  false,
		},
	}
	for _, test := range tests {
		if twoSum(test.numbers, test.target) != test.answer {
			t.Error("Expected", test.answer, "got", twoSum(test.numbers, test.target))
		}
	}
}
func ExampleTwoSum() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 7))
	// Output:
	// true
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{1, 2, 3, 4, 5, 6}, 7)
	}
}
