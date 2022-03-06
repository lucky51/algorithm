package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println(basicSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 21, 30, 90}, 7))
}
func TestLeftBound(t *testing.T) {
	fmt.Println(leftBound([]int{1, 2, 7, 7, 7, 7, 7, 8, 8, 20, 21, 30, 90}, 8))
}
