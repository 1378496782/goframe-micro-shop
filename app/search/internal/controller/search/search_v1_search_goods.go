package search

import (
	"context"
	"encoding/json"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"

	"shop-goframe-micro-service-refacotor/app/search/api/search/v1"
)

func (c *ControllerV1) SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error) {
	// 初始化响应结构
	response := &v1.SearchGoodsRes{
		List:  make([]*v1.GoodsInfoItem, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	// 调试：打印请求参数
	g.Log().Debugf(ctx, "Search request: %+v", req)

	// 错误类型
	infoError := consts.InfoError(consts.SearchGoods, consts.SearchFail)

	// 1. 获取ES客户端
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Errorf(ctx, "%v ES客户端未初始化", infoError)
		return nil, gerror.NewCode(gcode.CodeInternalError, "搜索服务暂不可用")
	}

	// 2. 构建查询条件
	boolQuery := elastic.NewBoolQuery()

	// 软删除过滤
	boolQuery.MustNot(elastic.NewWildcardQuery("deleted_at", "*?*"))

	if req.Keyword != "" {
		matchQuery := elastic.NewMatchQuery("name", req.Keyword)
		boolQuery.Must(matchQuery)
		g.Log().Debugf(ctx, "Adding keyword search: %s", req.Keyword)
	}

	if req.Brand != "" {
		termQuery := elastic.NewTermQuery("brand", req.Brand)
		boolQuery.Filter(termQuery)
		g.Log().Debugf(ctx, "Adding brand filter: %s", req.Brand)
	}

	if req.MinPrice > 0 || req.MaxPrice > 0 {
		rangeQuery := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			rangeQuery.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			rangeQuery.Lte(req.MaxPrice)
		}
		boolQuery.Filter(rangeQuery)
		g.Log().Debugf(ctx, "Adding price range: %d - %d", req.MinPrice, req.MaxPrice)
	}

	// 3. 构建搜索请求
	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	g.Log().Debugf(ctx, "Using index: %s", esIndexGoods)

	searchService := client.Search().Index(esIndexGoods).Query(boolQuery)
	searchService.From(int((req.Page - 1) * req.Size)).Size(int(req.Size))

	// 排序
	switch req.Sort {
	case "price_asc":
		searchService.Sort("price", true)
	case "price_desc":
		searchService.Sort("price", false)
	case "sale":
		searchService.Sort("sale", false)
	default:
		searchService.Sort("_score", false)
	}

	// 高亮
	highlight := elastic.NewHighlight().
		Field("name").
		PreTags("<em>").
		PostTags("</em>")
	searchService.Highlight(highlight)

	// 4. 执行搜索
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, infoError)
	}

	g.Log().Debugf(ctx, "Search result total hits: %d", searchResult.TotalHits())

	// 5. 处理结果
	response.Total = uint32(searchResult.TotalHits())

	for _, hit := range searchResult.Hits.Hits {
		g.Log().Debugf(ctx, "Processing hit: %s", hit.Id)

		var goods v1.GoodsInfoItem
		if err := json.Unmarshal(hit.Source, &goods); err != nil {
			g.Log().Errorf(ctx, "Failed to unmarshal hit: %v", err)
			continue
		}

		// 处理高亮
		if highlight, ok := hit.Highlight["name"]; ok && len(highlight) > 0 {
			goods.Highlight = highlight[0]
		} else {
			goods.Highlight = goods.Name
		}

		response.List = append(response.List, &goods)
	}

	g.Log().Debugf(ctx, "Returning %d items", len(response.List))
	return response, nil
}
