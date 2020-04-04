package Sub

import "testing"

func TestAverage(t *testing.T) {
	v1 := []int{1,2,3,4,5}
	v1Ans := Average(v1)

	if v1Ans != 3 {
		t.Error("extepcted 3 got a ", v1)
	}
}