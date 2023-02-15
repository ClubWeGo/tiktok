package command

import (
	"context"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest, fromID int64) ([]*comment.Comment, error) {
	Comments, err := db.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments, err := pack.Comments(s.ctx, Comments, fromID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
