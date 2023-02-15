package main

import (
	"context"
	"github.com/a76yyyy/tiktok/cmd/message/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	message "github.com/a76yyyy/tiktok/kitex_gen/message"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// MessageSrvImpl implements the last service interface defined in the IDL.
type MessageSrvImpl struct{}

// MessageAction implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildMessageActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.ActionType != 1 {
		resp = pack.BuildMessageActionResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewMessageActionService(ctx).MessageAction(req, claim.Id)
	if err != nil {
		resp = pack.BuildMessageActionResp(err)
		return resp, nil
	}
	resp = pack.BuildMessageActionResp(err)
	return resp, nil
}

// MessageChat implements the MessageSrvImpl interface.
func (s *MessageSrvImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageCharResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildMessageChatResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if claim.Id == 0 {
		resp = pack.BuildMessageChatResp(errno.ErrBind)
		return resp, nil
	}

	messages, err := command.NewMessageChatService(ctx).MessageChat(req, claim.Id)
	if err != nil {
		resp = pack.BuildMessageChatResp(err)
		return resp, nil
	}
	resp = pack.BuildMessageChatResp(errno.Success)
	resp.MessageList = messages
	//resp.MessageList = nil
	return resp, nil
}
