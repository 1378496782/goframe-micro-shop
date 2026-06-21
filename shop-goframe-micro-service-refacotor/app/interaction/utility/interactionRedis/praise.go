package interactionRedis

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	CommentLikeCountKeyPrefix = "interaction:comment:like_count:"
	commentLikeCountTTL       = int64(10 * 60)
)

func commentLikeCountKey(commentId uint32) string {
	return CommentLikeCountKeyPrefix + strconv.Itoa(int(commentId))
}

func GetCommentLikeCount(ctx context.Context, commentId uint32) (count uint32, ok bool, err error) {
	v, err := g.Redis().GroupString().Get(ctx, commentLikeCountKey(commentId))
	if err != nil {
		return 0, false, err
	}
	if v == nil || v.IsNil() || v.IsEmpty() || v.String() == "" || v.String() == "null" {
		return 0, false, nil
	}
	count = uint32(v.Int64())
	ok = true
	return
}

func SetCommentLikeCount(ctx context.Context, commentId uint32, count uint32) error {
	return g.Redis().GroupString().SetEX(ctx, commentLikeCountKey(commentId), count, commentLikeCountTTL)
}

func IncrCommentLikeCount(ctx context.Context, commentId uint32) error {
	key := commentLikeCountKey(commentId)
	if _, err := g.Redis().GroupString().Incr(ctx, key); err != nil {
		return err
	}
	return refreshCommentLikeCountTTL(ctx, key)
}

func DecrCommentLikeCount(ctx context.Context, commentId uint32) error {
	count, ok, err := GetCommentLikeCount(ctx, commentId)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	if count == 0 {
		return SetCommentLikeCount(ctx, commentId, 0)
	}

	key := commentLikeCountKey(commentId)
	next, err := g.Redis().GroupString().Decr(ctx, key)
	if err != nil {
		return err
	}
	if next < 0 {
		return SetCommentLikeCount(ctx, commentId, 0)
	}
	return refreshCommentLikeCountTTL(ctx, key)
}

func DeleteCommentLikeCount(ctx context.Context, commentId uint32) error {
	_, err := g.Redis().GroupGeneric().Del(ctx, commentLikeCountKey(commentId))
	return err
}

func refreshCommentLikeCountTTL(ctx context.Context, key string) error {
	_, err := g.Redis().GroupGeneric().Expire(ctx, key, commentLikeCountTTL)
	return err
}
