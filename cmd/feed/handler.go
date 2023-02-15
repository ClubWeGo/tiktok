package main

import (
	"context"

	"github.com/a76yyyy/tiktok/cmd/feed/command"
	"github.com/a76yyyy/tiktok/dal/pack"
	"github.com/a76yyyy/tiktok/kitex_gen/feed"
	"github.com/a76yyyy/tiktok/pkg/errno"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	var uid int64 = 0
	if *req.Token != "" {
		claim, err := Jwt.ParseToken(*req.Token)
		if err != nil {
			resp = pack.BuildVideoResp(errno.ErrTokenInvalid)
			return resp, nil
		} else {
			uid = claim.Id
		}
	}

	vis, nextTime, err := command.NewGetUserFeedService(ctx).GetUserFeed(req, uid)
	if err != nil {
		resp = pack.BuildVideoResp(err)
		return resp, nil
	}

	resp = pack.BuildVideoResp(errno.Success)
	resp.VideoList = vis
	resp.NextTime = &nextTime
	return resp, nil
}

// GetVideoById implements the FeedSrvImpl interface.
// Deprecated: Never Use GetVideoById
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	return
}
