package godiscogs

import "sort"
import "testing"

var splitTests = []struct {
	str string
	exp []string
}{
	{"abc123", []string{"abc", "123"}},
	{"IM 23", []string{"IM", "23"}},
}

func TestSplit(t *testing.T) {
	for _, tt := range splitTests {
		res := split(tt.str)
		if len(res) != len(tt.exp) {
			t.Errorf("Wrong length %v vs %v", res, tt.exp)
		} else {
			for i := range res {
				if res[i] != tt.exp[i] {
					t.Errorf("Mismatch %v vs %v", res, tt.exp)
				}
			}
		}
	}
}

var sortTests = []struct {
	r1 Release
	r2 Release
}{
	{Release{Labels: []*Label{&Label{Name: "TestOne"}}},
		Release{Labels: []*Label{&Label{Name: "TestTwo"}}}},

	{Release{Labels: []*Label{&Label{Name: "TestOne", Catno: "First"}}},
		Release{Labels: []*Label{&Label{Name: "TestOne", Catno: "Second"}}}},

	{Release{Labels: []*Label{&Label{Name: "TestOne", Catno: "IM 2"}}},
		Release{Labels: []*Label{&Label{Name: "TestOne", Catno: "IM 12"}}}},
}

var defaultComp = []struct {
	r1 Release
	r2 Release
}{
	{Release{Labels: []*Label{&Label{Name: "TestOne"}}},
		Release{Labels: []*Label{&Label{Name: "TestOne"}}}},
}

func TestFullSort(t *testing.T) {
	var releases []*Release
	releases = append(releases, &Release{Title: "First", Labels: []*Label{&Label{Name: "TestOne"}}})
	releases = append(releases, &Release{Title: "Last", Labels: []*Label{&Label{Name: "TestTwo"}}})
	releases = append(releases, &Release{Labels: []*Label{&Label{Name: "TestThree"}}})
	sort.Sort(ByLabelCat(releases))

	if releases[0].Title != "First" || releases[2].Title != "Last" {
		t.Errorf("Releases are not sorted correctly: %v", releases)
	}
}

func TestSortingByLabelCat(t *testing.T) {
	for _, tt := range sortTests {
		sValue := sortByLabelCat(tt.r1, tt.r2)
		if sValue >= 0 {
			t.Errorf("%v should come before %v (%v)", tt.r1, tt.r2, sValue)
		}
		sValueR := sortByLabelCat(tt.r2, tt.r1)
		if sValueR <= 0 {
			t.Errorf("%v should come before %v (%v)", tt.r1, tt.r2, sValueR)
		}
	}

	tt := defaultComp[0]
	sValue := sortByLabelCat(tt.r1, tt.r2)
	sValue2 := sortByLabelCat(tt.r2, tt.r1)
	if sValue != 0 || sValue2 != 0 {
		t.Errorf("Default is not zero: %v and %v", sValue, sValue2)
	}
}
