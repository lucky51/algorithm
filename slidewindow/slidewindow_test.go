package slidewindow

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	fmt.Println(checkInclusion("helloworld", "oll"))

}

func TestRespliceSlice(t *testing.T) {
	var s = "ADBECFEBANC"
	fmt.Println(s[7:10])
}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADBECFEBANC", "ABC"))
}
