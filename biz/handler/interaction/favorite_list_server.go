// Code generated by hertz generator.

package interaction

import (
	"context"
	"github.com/ClubWeGo/douyin/kitex_server"
	"github.com/ClubWeGo/douyin/tools"
	"github.com/ClubWeGo/douyin/tools/errno"

	interaction "github.com/ClubWeGo/douyin/biz/model/interaction"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteListMethod .
// @router /douyin/favorite/list/ [GET]
func FavoriteListMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interaction.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.NewErrNo(consts.StatusBadRequest, err.Error()), nil)
		return
	}
	valid, uid, err := tools.ValidateToken(req.Token)
	if err != nil {
		SendResponse(c, errno.NewErrNo(consts.StatusUnauthorized, err.Error()), nil)
		return
	}
	if !valid {
		SendResponse(c, errno.NewErrNo(consts.StatusUnauthorized, "token invalid"), nil)
		return
	}
	res, err := kitex_server.GetFavoriteList(ctx, uid)
	if err != nil {
		SendResponse(c, errno.RPCErr, nil)
		return
	}
	SendResponse(c, nil, res)
	return
}
