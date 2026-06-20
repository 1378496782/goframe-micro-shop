package praise_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	// PraiseTypeArticle 文章点赞
	PraiseTypeArticle = 1
	// PraiseTypeComment 评论/回复点赞
	PraiseTypeComment = 2
)

func Create(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	res = &v1.PraiseInfoCreateRes{}
	// 错误类型
	infoError := consts.InfoError(consts.PraiseInfo, consts.CreateFail)
	if req.Type == PraiseTypeComment {
		if err := ensureCommentLikeTarget(ctx, req.ObjectId); err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, err
		}
	}
	// 向数据库中插入数据并获取自动生成的ID
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		if isDuplicatePraiseError(err) {
			res.Id = uint32(id)
			return res, nil
		}
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	if req.Type == PraiseTypeArticle {
		// TODO文章点赞
	} else if req.Type == PraiseTypeComment {
		// 评论/回复点赞
		if err := increaseCommentLikeCount(ctx, req.ObjectId); err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
	}

	res.Id = uint32(id)
	return
}

func isDuplicatePraiseError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "Duplicate entry") || strings.Contains(errMsg, "Error 1062")
}

func Delete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error) {
	res = &v1.PraiseInfoDeleteRes{}
	// 根据ID从数据库中删除对应信息
	result, err := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		"id":        req.Id,
		"user_id":   req.UserId,
		"type":      req.Type,
		"object_id": req.ObjectId,
	}).Delete()
	infoError := consts.InfoError(consts.PraiseInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	res.Id = req.Id
	if rows == 0 {
		return
	}

	if req.Type == PraiseTypeArticle {
		// TODO文章点赞
	} else if req.Type == PraiseTypeComment {
		// 评论/回复点赞
		if err := decreaseCommentLikeCount(ctx, req.ObjectId); err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
	}

	return
}

func increaseCommentLikeCount(ctx context.Context, commentId uint32) error {
	_, err := dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().Id, commentId).
		WhereNull(dao.CommentInfo.Columns().DeletedAt).
		Increment(dao.CommentInfo.Columns().LikeCount, 1)
	return err
}

func ensureCommentLikeTarget(ctx context.Context, commentId uint32) error {
	count, err := dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().Id, commentId).
		WhereNull(dao.CommentInfo.Columns().DeletedAt).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, "评论不存在或已删除")
	}
	return nil
}

func decreaseCommentLikeCount(ctx context.Context, commentId uint32) error {
	_, err := dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().Id, commentId).
		WhereGT(dao.CommentInfo.Columns().LikeCount, 0).
		WhereNull(dao.CommentInfo.Columns().DeletedAt).
		Decrement(dao.CommentInfo.Columns().LikeCount, 1)
	return err
}
