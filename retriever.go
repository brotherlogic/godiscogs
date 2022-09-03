package godiscogs

import (
	"context"
	"fmt"
	"strconv"
)

type setRatingResponse struct {
	Username string
	Rating   int
	Message  string
}

// SetRating sets the rating on the specified releases
func (r *DiscogsRetriever) SetRating(ctx context.Context, releaseID int, rating int) error {
	if rating == 0 {
		r.delete(ctx, "/releases/"+strconv.Itoa(releaseID)+"/rating/BrotherLogic?token="+r.userToken, "")
		return nil
	}

	data, err := r.put(ctx, "/releases/"+strconv.Itoa(releaseID)+"/rating/BrotherLogic?token="+r.userToken, "{\"rating\": "+strconv.Itoa(rating)+"}")

	if err != nil {
		return err
	}

	var response setRatingResponse
	r.unmarshaller.Unmarshal(data, &response)

	if response.Rating == rating {
		return nil
	}

	return fmt.Errorf("Unable to rate release: %v", response.Message)
}
