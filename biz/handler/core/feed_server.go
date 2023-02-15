// Code generated by hertz generator.

package core

import (
	"context"
	"time"

	core "github.com/ClubWeGo/douyin/biz/model/core"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FeedMethod .
// @router /douyin/feed/ [GET]
func FeedMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	msgsucceed := "获取视频流成功"
	msgFailed := "获取视频流失败"

	resp := new(core.FeedResp)

	// 目前该api无需token，后续增加登录定制化内容则需根据token获取其他参数
	// var token string
	// if req.Token != nil { // 可选字段，需要验证是否存在，判断对应指针是否存在
	// 	token = *req.Token
	// }

	var latestTime = time.Now().UnixNano()
	if req.LatestTime != nil {
		latestTime = (*req.LatestTime) * 1e6 // app传入的是13位毫秒级时间戳，usermicro需传入纳秒级时间戳
	}
	r, err := kitex_server.Videoclient.GetVideosFeedMethod(context.Background(), &videomicro.GetVideosFeedReq{LatestTime: latestTime, Limit: 30})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &msgFailed
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.VideoList = make([]*core.Video, 0)
	for _, video := range r.VideoList {
		author, _ := kitex_server.GetUser(video.AuthorId)
		// 暂时不做处理，错误返回空对象即可
		resp.VideoList = append(resp.VideoList, &core.Video{
			ID:            video.Id,
			Author:        author,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false, // 需要增加喜欢配置
			Title:         video.Title,
		})
	}

	resp.StatusMsg = &msgsucceed
	nextTimeMs := (*r.NextTime) / 1e6 // 转为毫秒
	resp.NextTime = nextTimeMs

	c.JSON(consts.StatusOK, resp)
}
