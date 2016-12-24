package godiscogs

import (
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"testing"

	"github.com/golang/protobuf/proto"
)

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

	{Release{Title: "Low", Labels: []*Label{&Label{Name: "TestOne"}}},
		Release{Title: "VeryLow", Labels: []*Label{&Label{Name: "TestOne"}}}},

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

func TestFullSortWithAmbigousLabels(t *testing.T) {
	var releases []*Release
	releases = append(releases, &Release{Title: "First", Labels: []*Label{&Label{Name: "ToBeIgnore"}, &Label{Name: "TestOne"}}})
	releases = append(releases, &Release{Title: "Last", Labels: []*Label{&Label{Name: "ToBeIgnore"}, &Label{Name: "TestTwo"}}})
	releases = append(releases, &Release{Labels: []*Label{&Label{Name: "TestThree"}}})
	sort.Sort(ByLabelCat(releases))

	if releases[0].Title != "First" || releases[2].Title != "Last" {
		t.Errorf("Releases are not sorted correctly: %v", releases)
	}
}

func TestSortingOrderConsistency(t *testing.T) {
	data1, err := ioutil.ReadFile("testdata/sort_test/1/3139381.release")
	data2, _ := ioutil.ReadFile("testdata/sort_test/1/1531104.release")
	data3, _ := ioutil.ReadFile("testdata/sort_test/1/6512427.release")

	log.Printf("ERR %v", err)

	release1 := &Release{}
	release2 := &Release{}
	release3 := &Release{}
	proto.Unmarshal(data1, release1)
	proto.Unmarshal(data2, release2)
	proto.Unmarshal(data3, release3)

	log.Printf("WAH %v, %v", release1, data1)

	var cReleases []*Release
	cReleases = append(cReleases, release1)
	cReleases = append(cReleases, release2)
	cReleases = append(cReleases, release3)
	sort.Sort(ByLabelCat(cReleases))

	for i := 0; i < 100; i++ {
		var releases []*Release
		perm := rand.Perm(3)
		releases = append(releases, cReleases[perm[0]])
		releases = append(releases, cReleases[perm[1]])
		releases = append(releases, cReleases[perm[2]])
		sort.Sort(ByLabelCat(releases))

		failed := false
		for i := range releases {
			if releases[i].Id != cReleases[i].Id {
				failed = true
			}
		}

		if failed {
			t.Errorf("Sorting is not unique:")
			for j := range releases {
				t.Errorf("%v. %v -> %v", j, cReleases[j].Id, releases[j].Id)
			}
		}
	}
}

func TestSortingConsistencyWithMultipleLabels(t *testing.T) {
	r1 := Release{Labels: []*Label{&Label{Name: "REally Though", Catno: "IM 12"}, &Label{Name: "Actually First", Catno: "DDD"}}}
	r2 := Release{Labels: []*Label{&Label{Name: "TestOne", Catno: "IM 1"}, &Label{Name: "Behind", Catno: "BBB"}}}

	sValue := sortByLabelCat(r1, r2)
	log.Printf("HERE = %v", sValue)
	if sValue >= 0 {
		t.Errorf("Sorting is off with mulitple labels")
	}
	sValueR := sortByLabelCat(r2, r1)
	log.Printf("HERE = %v", sValueR)
	if sValueR <= 0 {
		t.Errorf("Reverse sorting is off with multiple labels")
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
