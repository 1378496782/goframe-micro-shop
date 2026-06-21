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

	v1 "shop-goframe-micro-service-refacotor/app/search/api/search/v1"
)

// SearchGoods 商品搜索接口。
//
// 基于 Elasticsearch 实现：按关键词全文匹配商品名，支持品牌、价格区间过滤，
// 支持按价格/销量/相关性排序，并对命中的商品名做关键词高亮。
// 整体流程：获取 ES 客户端 → 构建查询条件 → 组装搜索请求（分页/排序/高亮）
// → 执行搜索 → 反序列化命中结果并返回。
func (c *ControllerV1) SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error) {
	// 初始化响应结构：List 预置为空切片，避免无结果时返回 null
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

	// 2. 构建查询条件：用 bool query 组合多个子条件
	boolQuery := elastic.NewBoolQuery()

	// 软删除过滤：deleted_at 有任意非空值即视为已删除，用 wildcard "*?*"（至少一个字符）
	// 配合 MustNot 排除掉这些文档，只保留未删除商品。
	boolQuery.MustNot(elastic.NewWildcardQuery("deleted_at", "*?*"))

	// 关键词：用 match 做全文分词匹配商品名。放在 Must 里，会参与相关性打分（_score）。
	if req.Keyword != "" {
		// matchQuery := elastic.NewMatchQuery("name", req.Keyword)
		multiMatchQuery := elastic.NewMultiMatchQuery(
			req.Keyword,
			"name^4",
			"brand.text^3",
			"brand^3",
			"tags.text^2",
			"tags^2",
			"detail_info",
		).Type("best_fields")

		boolQuery.Must(multiMatchQuery)
		g.Log().Debugf(ctx, "Adding keyword search: %s", req.Keyword)
	}

	// 品牌：用 term 做精确匹配。放在 Filter 里，只做过滤、不算分，且可被 ES 缓存，性能更好。
	if req.Brand != "" {
		termQuery := elastic.NewTermQuery("brand", req.Brand)
		boolQuery.Filter(termQuery)
		g.Log().Debugf(ctx, "Adding brand filter: %s", req.Brand)
	}

	// 价格区间：min/max 任一 >0 才加 range 过滤；Gte/Lte 分别对应下界、上界（闭区间）。
	// 同样放 Filter 里，纯过滤不影响打分。
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
	// 分页：ES 的 From 是偏移量（从第几条开始），Size 是每页条数。
	// 第 page 页（page 从 1 开始）的偏移量 = (page-1) * size。
	searchService.From(int((req.Page - 1) * req.Size)).Size(int(req.Size))

	// 排序：Sort 第二个参数 ascending，true 升序 / false 降序。
	switch req.Sort {
	case "price_asc": // 价格升序
		searchService.Sort("price", true)
	case "price_desc": // 价格降序
		searchService.Sort("price", false)
	case "sale": // 按销量从高到低
		searchService.Sort("sale", false)
	default: // 默认按相关性得分 _score 从高到低
		searchService.Sort("_score", false)
	}

	// 高亮：对命中的商品名字段加 <em> 标签，前端可据此标红关键词
	highlight := elastic.NewHighlight().
		Field("name").
		Field("brand.text").
		Field("brand").
		Field("tags.text").
		Field("tags").
		Field("detail_info").
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
	// Total 是匹配的总数（用于前端分页），可能远大于本页返回的条数。
	response.Total = uint32(searchResult.TotalHits())

	// 遍历本页命中的文档，逐条反序列化成业务结构
	for _, hit := range searchResult.Hits.Hits {
		g.Log().Debugf(ctx, "Processing hit: %s", hit.Id)

		var goods v1.GoodsInfoItem
		// 把 ES 文档的原始 _source JSON 反序列化成商品对象；
		// 单条解析失败只跳过该条、不中断整批搜索。
		if err := json.Unmarshal(hit.Source, &goods); err != nil {
			g.Log().Errorf(ctx, "Failed to unmarshal hit: %v", err)
			continue
		}

		// 处理高亮：命中关键词时用 ES 返回的高亮片段（带 <em> 标签），
		// 否则回退用原始商品名，保证 Highlight 字段始终有值。
		goods.Highlight = pickHighlight(hit.Highlight, goods.Name)

		response.List = append(response.List, &goods)
	}

	g.Log().Debugf(ctx, "Returning %d items", len(response.List))
	return response, nil
}

func pickHighlight(highlights map[string][]string, fallback string) string {
	for _, field := range []string{"name", "brand.text", "brand", "tags.text", "tags", "detail_info"} {
		if values, ok := highlights[field]; ok && len(values) > 0 {
			return values[0]
		}
	}
	return fallback
}
