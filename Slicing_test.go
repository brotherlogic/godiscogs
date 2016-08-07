package godiscogs

import "log"
import "testing"

func TestSlice(t *testing.T) {
          var releases []*Release
          releases = append(releases, &Release{FormatQuantity: 2, Title: "First", Labels: []*Label{&Label{Name: "TestOne"}}})
	  releases = append(releases, &Release{FormatQuantity: 1, Title: "Last", Labels: []*Label{&Label{Name: "TestTwo"}}})
	  releases = append(releases, &Release{FormatQuantity: 1, Labels: []*Label{&Label{Name: "TestThree"}}})

	  splits := Split(releases,2)

	  log.Printf("SPLITS = %v", splits)

	  if len(splits[0]) != 1 {
	     t.Errorf("First split should have one entry: %v", splits)
	  }
	  if len(splits[1]) != 2 {
	     t.Errorf("Second split should have two entries: %v", splits)
	  }
}