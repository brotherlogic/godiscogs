package godiscogs

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func split(str string) []string {
	return regexp.MustCompile("[0-9]+|[a-z]+|[A-Z]+").FindAllString(str, -1)
}

// GetReleaseArtist Gets a string of the release artist of this record
func GetReleaseArtist(rel Release) string {
	if len(rel.Artists) > 0 {
		artistString := rel.Artists[0].Name
		for _, artist := range rel.Artists[1:] {
			artistString += " & " + artist.Name
		}
		return artistString
	}
	return "Unknown"
}

// ByLabelCat is a sorting function that sorts by label name, then catalogue number
type ByLabelCat []*Release

func (a ByLabelCat) Len() int           { return len(a) }
func (a ByLabelCat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLabelCat) Less(i, j int) bool { return sortByLabelCat(*a[i], *a[j]) < 0 }

func sortByLabelCat(rel1 Release, rel2 Release) int {
	label1 := GetMainLabel(rel1.Labels)
	label2 := GetMainLabel(rel2.Labels)

	labelSort := strings.Compare(label1.Name, label2.Name)
	if labelSort != 0 {
		return labelSort
	}

	cat1Elems := split(label1.Catno)
	cat2Elems := split(label2.Catno)

	toCheck := len(cat1Elems)
	if len(cat2Elems) < toCheck {
		toCheck = len(cat2Elems)
	}
	for i := 0; i < toCheck; i++ {
		if unicode.IsNumber(rune(cat1Elems[i][0])) && unicode.IsNumber(rune(cat2Elems[i][0])) {
			num1, _ := strconv.Atoi(cat1Elems[i])
			num2, _ := strconv.Atoi(cat2Elems[i])
			if num1 > num2 {
				return 1
			} else if num2 > num1 {
				return -1
			}
		} else {
			catComp := strings.Compare(cat1Elems[i], cat2Elems[i])
			if catComp != 0 {
				return catComp
			}
		}
	}

	//Fallout to sorting by title
	titleComp := strings.Compare(rel1.Title, rel2.Title)
	return titleComp
}

// GetMainLabel gets the main label from the release - this is the label to be used in e.g. sorting
func GetMainLabel(labels []*Label) *Label {
	if len(labels) == 0 {
		return nil
	} else if len(labels) == 1 {
		return labels[0]
	} else {
		labelName := labels[0].Name
		labelCat := labels[0].Catno
		labelIndex := 0

		for i, label := range labels[1:] {
			if strings.Compare(labelName, label.Name) > 0 || (strings.Compare(labelName, label.Name) == 0 && strings.Compare(labelCat, label.Catno) > 0) {
				labelName = label.Name
				labelCat = label.Catno
				labelIndex = i + 1
			}
		}

		return labels[labelIndex]
	}
}

// Split splits a releases list into buckets
func Split(releases []*Release, n float64) [][]*Release {
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
		if currentValue+float64(rel.FormatQuantity) > boundaryValue {
			solution = append(solution, currentReleases)
			currentReleases = make([]*Release, 0)
			boundaryValue += boundaryAccumulator
		}

		currentReleases = append(currentReleases, rel)
		currentValue += float64(rel.FormatQuantity)
	}
	solution = append(solution, currentReleases)

	return solution
}
