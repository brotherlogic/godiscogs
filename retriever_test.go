package godiscogs

import "testing"

func TestSetRating(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	err := retr.SetRating(10567529, 5)
	if err != nil {
		t.Errorf("Rating has not been set")
	}
}

func TestNullRating(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	err := retr.SetRating(10567529, 0)
	if err != nil {
		t.Errorf("Rating has not been set")
	}
}

func TestSetRatingFail(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	err := retr.SetRating(10567528, 5)
	if err == nil {
		t.Errorf("Fail set rating has not failed")
	}
}

func TestSetRatingPutFail(t *testing.T) {
	retr := NewTestDiscogsRetriever()
	retr.getter = testFailGetter{}
	err := retr.SetRating(2000000000, 5)
	if err == nil {
		t.Errorf("Fail set rating has not failed")
	}
}
