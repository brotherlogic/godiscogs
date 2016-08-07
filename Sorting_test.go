package godiscogs

import "sort"
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
	for _, tt := range sortTests[:len(sortTests)-1] {
		sValue := sortByLabelCat(tt.r1, tt.r2)
		if sValue >= 0 {
			t.Errorf("%v should come before %v", tt.r1, tt.r2)
		}
		sValueR := sortByLabelCat(tt.r2, tt.r1)
		if sValueR <= 0 {
			t.Errorf("%v should come before %v", tt.r1, tt.r2)
		}
	}

	tt := sortTests[len(sortTests)-1]
	sValue := sortByLabelCat(tt.r1, tt.r2)
	sValue2 := sortByLabelCat(tt.r2, tt.r1)
	if sValue != 0 || sValue2 != 0 {
		t.Errorf("Default is not zero: %v and %v", sValue, sValue2)
	}
}
