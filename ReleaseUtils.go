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

// Sort sorts two releases into the order we like
func Sort(rel1 Release, rel2 Release) int {
	//First sort by label - just use the first one in the list
	label1 := rel1.Labels[0]
	label2 := rel2.Labels[0]

	labelSort := strings.Compare(label1.Name, label2.Name)
	if labelSort != 0 {
		return labelSort
	}

	return 0
}
