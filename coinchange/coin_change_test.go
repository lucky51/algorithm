package coinchange

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5, 6, 10}, 20))
}

func TestRuneEquals(t *testing.T) {

	r := rune('1')

	fmt.Println(string((r+1)%48 + 48))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r+1)%48 + 48
	fmt.Println(string(r))
	r = (r-1)%48 + 48
	fmt.Println(string(r))
	r = (r-1)%48 + 48
	fmt.Println(string(r))
}
