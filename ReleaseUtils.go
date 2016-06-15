package godiscogs

// GetReleaseArtist Gets a string of the release artist of this record
func GetReleaseArtist(rel Release) string {
	artistString := rel.Artists[0].Name
	for _, artist := range rel.Artists[1:] {
		artistString += " & " + artist.Name
	}
	return artistString
}
