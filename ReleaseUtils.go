package godiscogs

import "strings"

// GetReleaseArtist Gets a string of the release artist of this record
func GetReleaseArtist(rel Release) string {
	artistString := rel.Artists[0].Name
	for _, artist := range rel.Artists[1:] {
		artistString += " & " + artist.Name
	}
	return artistString
}

// ByLabelCat is a sorting function that sorts by label name, then catalogue number
type ByLabelCat []*Release

func (a ByLabelCat) Len() int           { return len(a) }
func (a ByLabelCat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLabelCat) Less(i, j int) bool { return sortByLabelCat(*a[i], *a[j]) < 0 }

func sortByLabelCat(rel1 Release, rel2 Release) int {
	label1 := rel1.Labels[0]
	label2 := rel2.Labels[0]

	labelSort := strings.Compare(label1.Name, label2.Name)
	if labelSort != 0 {
		return labelSort
	}

	return 0
}

// Split splits a releases list into buckets
func Split(releases []*Release, n float64) [][]*Release{
     var solution [][]*Release

     var count int32
     count = 0
     for _, rel := range releases {
     	 count += rel.FormatQuantity
     }

     boundaryAccumulator := float64(count) / n
     boundaryValue := boundaryAccumulator
     currentValue := 0.0
     var currentReleases []*Release
     for _, rel := range releases {
     	 if currentValue + float64(rel.FormatQuantity) > boundaryValue {
	    solution = append(solution, currentReleases)
	    currentReleases = make([]*Release, 0)
	    currentValue = 0.0
	    boundaryValue += boundaryAccumulator
	 }

	 currentReleases = append(currentReleases, rel)
	 currentValue += float64(rel.FormatQuantity)
     }
	    solution = append(solution, currentReleases)
	    
     return solution
}