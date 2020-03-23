package twoSum_1

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type example struct {
	nums   []int
	target int
	result []int
}

var examples []example

func init() {
	examples = append(examples, example{[]int{2, 7, 5, 3, 4}, 9, []int{0, 1}})
	examples = append(examples, example{[]int{5, 6, 8, 3, 1, 2, 54}, 41, []int{-1, -1}})
	examples = append(examples, example{[]int{3, 5, 2, 4, 9, 7, 2}, 13, []int{3, 4}})
	examples = append(examples, example{[]int{5, 6, 8, 3, 1, 2, 5}, 6, []int{0, 4}})
	examples = append(examples, example{[]int{5, 6, 18, 3, 1, 2, 54}, 24, []int{1, 2}})
}

func TestTwoSum(t *testing.T) {
	for _, example := range examples {
		if !reflect.DeepEqual(TwoSum(example.nums, example.target), example.result) {
			fmt.Println(TwoSum(example.nums, example.target), example.result)
			t.Error("error")
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := i % 4
		TwoSum(examples[i].nums, examples[i].target)
	}
}

func BenchmarkTwoSumParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := rand.Intn(111) % 4
			TwoSum(examples[i].nums, examples[i].target)
		}
	})
}
