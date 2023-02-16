// Code generated by hertz generator.

package relation

import (
	"context"

	relation "github.com/ClubWeGo/douyin/biz/model/relation"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RelationMethod .
// @router /douyin/relation/action/ [POST]
func RelationMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.RelationReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.RelationResp)

	c.JSON(consts.StatusOK, resp)
}
