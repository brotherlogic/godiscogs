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
