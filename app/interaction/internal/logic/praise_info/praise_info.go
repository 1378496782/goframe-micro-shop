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

func Create(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	res = &v1.PraiseInfoCreateRes{}
	// 错误类型
	infoError := consts.InfoError(consts.PraiseInfo, consts.CreateFail)
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
	// 根据ID从数据库中删除对应信息
	_, err = dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		"id":      req.Id,
		"user_id": req.UserId,
	}).Delete()
	infoError := consts.InfoError(consts.PraiseInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	// rows, err := result.RowsAffected()
	// if err != nil {
	// 	g.Log().Errorf(ctx, "%v %v", infoError, err)
	// 	return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	// }
	// if rows == 0 {
	// return nil, gerror.New("点赞不存在")
	// }

	// 返回删除成功的空响应
	return &v1.PraiseInfoDeleteRes{Id: req.Id}, nil // 返回空结构体
}
