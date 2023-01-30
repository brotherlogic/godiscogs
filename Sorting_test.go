package godiscogs

import (
	"io/ioutil"
	"math/rand"
	"sort"
	"testing"

	pb "github.com/brotherlogic/godiscogs/proto"

	"google.golang.org/protobuf/proto"
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
	r1 pb.Release
	r2 pb.Release
}{
	{pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}},
		pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestTwo"}}}},

	{pb.Release{Title: "Low", Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}},
		pb.Release{Title: "VeryLow", Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}}},

	{pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne", Catno: "First"}}},
		pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne", Catno: "Second"}}}},

	{pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne", Catno: "IM 2"}}},
		pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne", Catno: "IM 12"}}}},
}

var defaultComp = []struct {
	r1 pb.Release
	r2 pb.Release
}{
	{pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}},
		pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}}},
}

func TestFullSort(t *testing.T) {
	var Releases []*pb.Release
	Releases = append(Releases, &pb.Release{Title: "First", Labels: []*pb.Label{&pb.Label{Name: "TestOne"}}})
	Releases = append(Releases, &pb.Release{Title: "Last", Labels: []*pb.Label{&pb.Label{Name: "TestTwo"}}})
	Releases = append(pb.Releases, &pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestThree"}}})
	sort.Sort(ByLabelCat(pb.Releases))

	if pb.Releases[0].Title != "First" || pb.Releases[2].Title != "Last" {
		t.Errorf("pb.Releases are not sorted correctly: %v", Releases)
	}
}

func TestFullSortWithAmbigousLabels(t *testing.T) {
	var Releases []*pb.Release
	Releases = append(Releases, &pb.Release{Title: "First", Labels: []*pb.Label{&pb.Label{Name: "ToBeIgnore"}, &pb.Label{Name: "TestOne"}}})
	Releases = append(Releases, &pb.Release{Title: "Last", Labels: []*pb.Label{&pb.Label{Name: "ToBeIgnore"}, &pb.Label{Name: "TestTwo"}}})
	Releases = append(Releases, &pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestThree"}}})
	sort.Sort(ByLabelCat(Releases))

	if pb.Releases[0].Title != "First" || pb.Releases[2].Title != "Last" {
		t.Errorf("pb.Releases are not sorted correctly: %v", Releases)
	}
}

func TestSortingOrderConsistency(t *testing.T) {
	data1, _ := ioutil.ReadFile("testdata/sort_test/1/3139381.pb.Release")
	data2, _ := ioutil.ReadFile("testdata/sort_test/1/1531104.pb.Release")
	data3, _ := ioutil.ReadFile("testdata/sort_test/1/6512427.pb.Release")

	Release1 := &pb.Release{}
	Release2 := &pb.Release{}
	Release3 := &pb.Release{}
	proto.Unmarshal(data1, pb.Release1)
	proto.Unmarshal(data2, pb.Release2)
	proto.Unmarshal(data3, pb.Release3)

	var cReleases []*pb.Release
	cReleases = append(cReleases, Release1)
	cReleases = append(cReleases, Release2)
	cReleases = append(cReleases, Release3)
	sort.Sort(ByLabelCat(cReleases))

	for i := 0; i < 100; i++ {
		var Releases []*pb.Release
		perm := rand.Perm(3)
		Releases = append(Releases, cReleases[perm[0]])
		Releases = append(Releases, cReleases[perm[1]])
		Releases = append(Releases, cReleases[perm[2]])
		sort.Sort(ByLabelCat(pb.Releases))

		failed := false
		for i := range pb.Releases {
			if Releases[i].Id != cReleases[i].Id {
				failed = true
			}
		}

		if failed {
			t.Errorf("Sorting is not unique:")
			for j := range pb.Releases {
				t.Errorf("%v. %v -> %v", j, cReleases[j].Id, Releases[j].Id)
			}
		}
	}
}

func TestSortingConsistencyWithMultipleLabels(t *testing.T) {
	r1 := pb.Release{Labels: []*pb.Label{&pb.Label{Name: "REally Though", Catno: "IM 12"}, &pb.Label{Name: "Actually First", Catno: "DDD"}}}
	r2 := pb.Release{Labels: []*pb.Label{&pb.Label{Name: "TestOne", Catno: "IM 1"}, &pb.Label{Name: "Behind", Catno: "BBB"}}}

	sValue := sortByLabelCat(r1, r2)
	if sValue >= 0 {
		t.Errorf("Sorting is off with mulitple Labels")
	}
	sValueR := sortByLabelCat(r2, r1)
	if sValueR <= 0 {
		t.Errorf("Reverse sorting is off with multiple Labels")
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
