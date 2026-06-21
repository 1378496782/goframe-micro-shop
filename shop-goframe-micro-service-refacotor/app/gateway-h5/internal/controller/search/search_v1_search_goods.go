package search

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/search/v1"
)

type searchServiceResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    *v1.SearchGoodsRes `json:"data"`
}

func (c *ControllerV1) SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error) {
	serviceURL := strings.TrimRight(
		g.Cfg().MustGet(ctx, "search.serviceUrl", "http://search-service:8499").String(),
		"/",
	)

	httpRes, err := g.Client().
		Timeout(5*time.Second).
		Get(ctx, serviceURL+"/search/goods", g.Map{
			"keyword":   req.Keyword,
			"brand":     req.Brand,
			"min_price": req.MinPrice,
			"max_price": req.MaxPrice,
			"sort":      req.Sort,
			"page":      req.Page,
			"size":      req.Size,
		})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "搜索服务调用失败")
	}
	defer httpRes.Close()

	body := httpRes.ReadAll()
	if httpRes.StatusCode != http.StatusOK {
		return nil, gerror.NewCodef(
			gcode.CodeInternalError,
			"搜索服务返回异常: status=%d body=%s",
			httpRes.StatusCode,
			string(body),
		)
	}

	var upstream searchServiceResponse
	if err := gjson.DecodeTo(body, &upstream); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "搜索服务响应解析失败")
	}
	if upstream.Code != 0 {
		return nil, gerror.NewCodef(gcode.CodeInternalError, "搜索服务返回失败: %s", upstream.Message)
	}
	if upstream.Data == nil {
		return &v1.SearchGoodsRes{
			List: make([]*v1.GoodsInfoItem, 0),
			Page: req.Page,
			Size: req.Size,
		}, nil
	}

	return upstream.Data, nil
}
