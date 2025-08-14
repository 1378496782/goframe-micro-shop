package consts

const (
	CollectionInfo = "CollectionInfo"
	CommentInfo    = "CommentInfo"
	PraiseInfo     = "PraiseInfo"
	GetListFail    = "GetList 查询失败"
	CreateFail     = "Create 插入失败"
	DeleteFail     = "Delete 删除失败"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
