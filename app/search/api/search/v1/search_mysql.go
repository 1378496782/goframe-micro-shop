package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SearchGoodsMysqlReq struct {
	g.Meta   `path:"/search/goods/mysql" method:"get" tags:"搜索" sm:"商品数据库搜索"`
	Keyword  string `json:"keyword" dc:"搜索关键词"`
	Brand    string `json:"brand" dc:"品牌名称"`
	MinPrice uint64 `json:"min_price" dc:"最低价格(分)"`
	MaxPrice uint64 `json:"max_price" dc:"最高价格(分)"`
	Sort     string `json:"sort" dc:"排序方式: default-默认, price_asc-价格升序, price_desc-价格降序, sale-销量"`
	Page     uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size     uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type SearchGoodsMysqlRes struct {
	List  []*GoodsInfoItem `json:"list" dc:"商品列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}
