package search

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"shop-goframe-micro-service-refacotor/app/search/api/search/v1"
	"shop-goframe-micro-service-refacotor/app/search/internal/dao"
	"shop-goframe-micro-service-refacotor/app/search/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/consts"
)

func (c *ControllerV1) SearchGoodsMysql(ctx context.Context, req *v1.SearchGoodsMysqlReq) (res *v1.SearchGoodsMysqlRes, err error) {
	response := &v1.SearchGoodsMysqlRes{
		List:  make([]*v1.GoodsInfoItem, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	// 错误类型
	infoError := consts.InfoError(consts.SearchGoods, consts.SearchFail)

	// 构建查询
	query := dao.GoodsInfo.Ctx(ctx).Where("deleted_at IS NULL")

	// 关键词（名称包含）
	if req.Keyword != "" {
		query = query.WhereLike("name", fmt.Sprintf("%%%s%%", req.Keyword))
	}

	// 品牌过滤
	if req.Brand != "" {
		query = query.Where("brand", req.Brand)
	}

	// 价格区间
	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}
	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	// 总数
	total, err := query.Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 排序
	switch req.Sort {
	case "price_asc":
		query = query.OrderAsc("price")
	case "price_desc":
		query = query.OrderDesc("price")
	case "sale":
		query = query.OrderDesc("sale")
	default:
		// 默认按 sort 降序，再按 created_at 降序
		query = query.OrderDesc("sort").OrderDesc("created_at")
	}

	// 分页查询
	records, err := query.
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 转换
	for _, record := range records {
		var e entity.GoodsInfo
		if err := record.Struct(&e); err != nil {
			continue
		}
		var item v1.GoodsInfoItem
		if err := gconv.Struct(e, &item); err != nil {
			continue
		}
		// 时间字段处理为字符串
		item.CreatedAt = toTimeString(e.CreatedAt)
		item.UpdatedAt = toTimeString(e.UpdatedAt)
		item.DeletedAt = toTimeString(e.DeletedAt)
		item.Highlight = item.Name

		response.List = append(response.List, &item)
	}

	return response, nil
}
func toTimeString(t *gtime.Time) string {
	if t == nil {
		return ""
	}
	return t.String()
}
