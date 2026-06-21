package comment_info

import (
	"context"

	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/interaction/utility/interactionRedis"

	"github.com/gogf/gf/v2/frame/g"
)

const praiseTypeComment = 2

type LikeCountCalibrateResult struct {
	Scanned uint32
	Fixed   uint32
}

type praiseCountRow struct {
	ObjectId int `json:"object_id"`
	Count    int `json:"count"`
}

func CalibrateCommentLikeCount(ctx context.Context, batchSize int) (*LikeCountCalibrateResult, error) {
	if batchSize <= 0 {
		batchSize = 500
	}

	result := &LikeCountCalibrateResult{}
	lastId := 0

	for {
		var comments []entity.CommentInfo
		if err := dao.CommentInfo.Ctx(ctx).
			Fields(
				dao.CommentInfo.Columns().Id,
				dao.CommentInfo.Columns().LikeCount,
			).
			WhereGT(dao.CommentInfo.Columns().Id, lastId).
			WhereNull(dao.CommentInfo.Columns().DeletedAt).
			OrderAsc(dao.CommentInfo.Columns().Id).
			Limit(batchSize).
			Scan(&comments); err != nil {
			return result, err
		}
		if len(comments) == 0 {
			break
		}

		commentIds := make([]int, 0, len(comments))
		for _, comment := range comments {
			commentIds = append(commentIds, comment.Id)
			if comment.Id > lastId {
				lastId = comment.Id
			}
		}

		praiseCountMap, err := getPraiseCountMap(ctx, commentIds)
		if err != nil {
			return result, err
		}

		for _, comment := range comments {
			actualCount := praiseCountMap[comment.Id]
			result.Scanned++

			if comment.LikeCount != actualCount {
				if _, err := dao.CommentInfo.Ctx(ctx).
					Where(dao.CommentInfo.Columns().Id, comment.Id).
					WhereNull(dao.CommentInfo.Columns().DeletedAt).
					Data(g.Map{
						dao.CommentInfo.Columns().LikeCount: actualCount,
					}).
					Update(); err != nil {
					return result, err
				}
				result.Fixed++
			}

			if err := interactionRedis.SetCommentLikeCount(ctx, uint32(comment.Id), uint32(actualCount)); err != nil {
				g.Log().Warningf(ctx, "校准评论点赞 Redis 计数失败: comment_id=%d, err=%v", comment.Id, err)
			}
		}
	}

	return result, nil
}

func getPraiseCountMap(ctx context.Context, commentIds []int) (map[int]int, error) {
	countMap := make(map[int]int, len(commentIds))
	if len(commentIds) == 0 {
		return countMap, nil
	}

	var rows []praiseCountRow
	if err := dao.PraiseInfo.Ctx(ctx).
		Fields("object_id, COUNT(*) as count").
		Where(dao.PraiseInfo.Columns().Type, praiseTypeComment).
		WhereIn(dao.PraiseInfo.Columns().ObjectId, commentIds).
		Group(dao.PraiseInfo.Columns().ObjectId).
		Scan(&rows); err != nil {
		return nil, err
	}

	for _, row := range rows {
		countMap[row.ObjectId] = row.Count
	}

	return countMap, nil
}
