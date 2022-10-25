package godiscogs

import (
	"context"
	"testing"
)

func TestGetArtist(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	release, _ := retr.GetRelease(context.Background(), 6099374)
	artist := GetReleaseArtist(release)

	if artist != "Bill Comeau & Pete Levin" {
		t.Errorf("Artist is incorrect: %v", artist)
	}
}
