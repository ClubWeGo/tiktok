package command

import (
	"context"
	"github.com/a76yyyy/tiktok/kitex_gen/message"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

type MessageActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

// CommentActionService action comment.
func (s *MessageActionService) MessageAction(req *message.DouyinMessageActionRequest, uid int64) error {
	// 1-评论
	if req.ActionType == 1 {
		return db.NewMessage(s.ctx, &db.Message{
			UserID:   int(uid),
			ToUserID: int(req.ToUserId),
			Content:  req.Content,
		})
	}
	return errno.ErrBind
}
