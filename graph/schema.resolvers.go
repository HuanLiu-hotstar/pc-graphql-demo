package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pc-compositor/graph/data"
	"pc-compositor/graph/generated"
	"pc-compositor/graph/model"
)

func (r *queryResolver) GetPlaybackUrls(ctx context.Context, input *model.Input) (model.PlaybackResult, error) {
	p := data.GetPlaybacks(input.ContentID)
	if p == nil {
		failure := model.FailureResult{
			ErrCode: 1,
			ErrMsg:  fmt.Sprintf("not found:%s ", input.ContentID),
		}
		return failure, nil
	}
	success := model.SuccessResult{
		Playback: p,
	}
	return success, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
