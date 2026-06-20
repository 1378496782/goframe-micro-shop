package interaction

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	comment "shop-goframe-micro-service-refacotor/app/interaction/api/comment_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func convertCommentInfoItem(comment *pbentity.CommentInfo) *v1.CommentInfoItem {
	if comment == nil {
		return nil
	}
	item := &v1.CommentInfoItem{
		Id:        uint32(comment.Id),
		UserId:    uint32(comment.UserId),
		ObjectId:  uint32(comment.ObjectId),
		Type:      uint32(comment.Type),
		ParentId:  uint32(comment.ParentId),
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Replies:   make([]*v1.CommentInfoItem, 0, len(comment.Replies)),
	}
	for _, reply := range comment.Replies {
		item.Replies = append(item.Replies, convertCommentInfoItem(reply))
	}
	return item
}

func (c *ControllerV1) CommentInfoGetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &comment.CommentInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.CommentInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.CommentInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
		List:  make([]*v1.CommentInfoItem, 0, len(grpcRes.Data.List)),
	}

	for _, comment := range grpcRes.Data.List {
		res.List = append(res.List, convertCommentInfoItem(comment))
	}

	return res, nil
}
