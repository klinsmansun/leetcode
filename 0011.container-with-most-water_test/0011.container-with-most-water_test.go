package leetcode

import (
	"fmt"
	"testing"
)

var MostWaterData []int

func TestContainerWithMostWater(t *testing.T) {
	result := ContainerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Fail()
	}
	result = ContainerWithMostWater([]int{1, 2, 4, 3})
	if result != 4 {
		fmt.Println("expect 4, got ", result)
		t.Fail()
	}
}

func BenchmarkContainerWithMostWater(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		ContainerWithMostWater(MostWaterData)
	}
}

func init() {
	MostWaterData = []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 12, 6, 9}
}
