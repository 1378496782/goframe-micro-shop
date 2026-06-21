package praise_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/praise_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/utility/interactionRedis"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
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

	var id int64
	var praiseChanged bool
	txErr := dao.PraiseInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if req.Type == PraiseTypeComment {
			if err := ensureCommentLikeTargetTx(ctx, tx, req.ObjectId); err != nil {
				g.Log().Errorf(ctx, "%v %v", infoError, err)
				return err
			}
		}
		// 向数据库中插入数据并获取自动生成的ID
		id, err = tx.Model(dao.PraiseInfo.Table()).InsertAndGetId(req)
		if err != nil {
			if isDuplicatePraiseError(err) {
				res.Id = uint32(id)
				return nil
			}
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}

		if req.Type == PraiseTypeArticle {
			// TODO文章点赞
		} else if req.Type == PraiseTypeComment {
			// 评论/回复点赞
			if err := increaseCommentLikeCountTx(ctx, tx, req.ObjectId); err != nil {
				g.Log().Errorf(ctx, "%v %v", infoError, err)
				return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
			}
		}
		praiseChanged = true
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	// 增加点赞数
	if praiseChanged && req.Type == PraiseTypeComment {
		if cacheErr := interactionRedis.IncrCommentLikeCount(ctx, req.ObjectId); cacheErr != nil {
			g.Log().Warningf(ctx, "增加评论点赞 Redis 计数失败: comment_id=%d, err=%v", req.ObjectId, cacheErr)
			_ = interactionRedis.DeleteCommentLikeCount(ctx, req.ObjectId)
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
	var praiseChanged bool
	txErr := dao.PraiseInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Model(dao.PraiseInfo.Table()).
			Where(g.Map{
				"id":        req.Id,
				"user_id":   req.UserId,
				"type":      req.Type,
				"object_id": req.ObjectId,
			}).Delete()
		infoError := consts.InfoError(consts.PraiseInfo, consts.DeleteFail)
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
		rows, err := result.RowsAffected()
		if err != nil {
			g.Log().Errorf(ctx, "%v %v", infoError, err)
			return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
		}
		res.Id = req.Id
		if rows == 0 {
			return nil
		}

		if req.Type == PraiseTypeArticle {
			// TODO文章点赞
		} else if req.Type == PraiseTypeComment {
			// 评论/回复点赞
			if err := decreaseCommentLikeCountTx(ctx, tx, req.ObjectId); err != nil {
				g.Log().Errorf(ctx, "%v %v", infoError, err)
				return gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
			}
		}

		praiseChanged = true
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	if praiseChanged && req.Type == PraiseTypeComment {
		if cacheErr := interactionRedis.DecrCommentLikeCount(ctx, req.ObjectId); cacheErr != nil {
			g.Log().Warningf(ctx, "减少评论点赞 Redis 计数失败: comment_id=%d, err=%v", req.ObjectId, cacheErr)
			_ = interactionRedis.DeleteCommentLikeCount(ctx, req.ObjectId)
		}
	}
	return
}

func increaseCommentLikeCountTx(ctx context.Context, tx gdb.TX, commentId uint32) error {
	_, err := tx.Model(dao.CommentInfo.Table()).
		Where(dao.CommentInfo.Columns().Id, commentId).
		WhereNull(dao.CommentInfo.Columns().DeletedAt).
		Increment(dao.CommentInfo.Columns().LikeCount, 1)
	return err
}

func ensureCommentLikeTargetTx(ctx context.Context, tx gdb.TX, commentId uint32) error {
	count, err := tx.Model(dao.CommentInfo.Table()).
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

func decreaseCommentLikeCountTx(ctx context.Context, tx gdb.TX, commentId uint32) error {
	_, err := tx.Model(dao.CommentInfo.Table()).
		Where(dao.CommentInfo.Columns().Id, commentId).
		WhereGT(dao.CommentInfo.Columns().LikeCount, 0).
		WhereNull(dao.CommentInfo.Columns().DeletedAt).
		Decrement(dao.CommentInfo.Columns().LikeCount, 1)
	return err
}
