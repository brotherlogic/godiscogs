package godiscogs

import "testing"

var sortTests = []struct {
	r1 Release
	r2 Release
}{
	{Release{Labels: []*Label{&Label{Name: "TestOne"}}},
		Release{Labels: []*Label{&Label{Name: "TestTwo"}}}},

	{Release{Labels: []*Label{&Label{Name: "TestOne"}}},
		Release{Labels: []*Label{&Label{Name: "TestOne"}}}},
}

func TestSorting(t *testing.T) {
	for _, tt := range sortTests[:len(sortTests)-1] {
		sValue := Sort(tt.r1, tt.r2)
		if sValue >= 0 {
			t.Errorf("%v should come before %v", tt.r1, tt.r2)
		}
		sValueR := Sort(tt.r2, tt.r1)
		if sValueR <= 0 {
			t.Errorf("%v should come before %v", tt.r1, tt.r2)
		}
	}

	tt := sortTests[len(sortTests)-1]
	sValue := Sort(tt.r1, tt.r2)
	sValue2 := Sort(tt.r2, tt.r1)
	if sValue != 0 || sValue2 != 0 {
		t.Errorf("Default is not zero: %v and %v", sValue, sValue2)
	}
}
