package slidewindow

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	fmt.Println(checkInclusion("eidboaoo", "ab"))
}

func TestRespliceSlice(t *testing.T) {
	var s = "ADBECFEBANC"
	fmt.Println(s[7:10])
}

func TestMinWindow(t *testing.T) {

	fmt.Println(minWindow("ADBECFEBANC", "ABC"))
	fmt.Println(minWindow1("ADBECFEBANC", "ABC"))
}

func TestFindAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("baa", "aa"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}
